package usecasePhoto

import (
	"context"

	"github.com/jinzhu/copier"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoPhoto "github.com/madmuzz05/go-final-project/service/photo/dto"
	entityPhoto "github.com/madmuzz05/go-final-project/service/photo/entity"
	entityUser "github.com/madmuzz05/go-final-project/service/user/entity"
)

func (u *PhotoUsecase) GetOne(ctx context.Context, id int) (res dtoPhoto.PhotoResponse, err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	model, errModel := u.PhotoRepository.GetOne(ctx, id)
	if errModel != nil {
		err = errModel
		return
	}
	var reqUser = entityUser.User{}
	reqUser.Id = model.UserId
	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	copier.Copy(&res, &model)
	if errUser == nil {
		res.User = &user
	}

	comments, errComments := u.CommentRepository.GetCommentByPhotoId(ctx, int(model.Id))
	if errComments == nil {
		res.Comment = &comments
	}
	return
}

func (u *PhotoUsecase) GetAll(ctx context.Context) (res []dtoPhoto.PhotoResponse, err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	model, errModel := u.PhotoRepository.GetAll(ctx)
	if errModel != nil {
		err = errModel
		return
	}

	for _, v := range model {
		var tempRes dtoPhoto.PhotoResponse
		copier.Copy(&tempRes, &v)

		var reqUser = entityUser.User{}
		reqUser.Id = v.UserId
		user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
		if errUser == nil {
			tempRes.User = &user
		}
		comments, errComments := u.CommentRepository.GetCommentByPhotoId(ctx, int(v.Id))
		if errComments == nil {
			tempRes.Comment = &comments
		}

		res = append(res, tempRes)
	}
	return
}

func (u *PhotoUsecase) CreatePhoto(ctx context.Context, req entityPhoto.Photo) (res dtoPhoto.PhotoResponse, err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	var reqUser = entityUser.User{}
	reqUser.Id = req.UserId

	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	if errUser != nil {
		err = errUser
		return
	}

	model, errModel := u.PhotoRepository.CreatePhoto(ctx, req)
	if errModel != nil {
		err = errModel
		return
	}

	copier.Copy(&res, &model)
	res.User = &user
	return
}
func (u *PhotoUsecase) UpdatePhoto(ctx context.Context, req entityPhoto.Photo, id int) (res dtoPhoto.PhotoResponse, err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	model, errModel := u.PhotoRepository.UpdatePhoto(ctx, req, id)
	if errModel != nil {
		err = errModel
		return
	}
	var reqUser = entityUser.User{}
	reqUser.Id = model.UserId
	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	copier.Copy(&res, &model)
	if errUser == nil {
		res.User = &user
	}
	return
}

func (u *PhotoUsecase) DeletePhoto(ctx context.Context, id int) (err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	errModel := u.PhotoRepository.DeletePhoto(ctx, id)
	if errModel != nil {
		err = errModel
		return
	}
	return
}
