package usecaseComment

import (
	"context"

	"github.com/jinzhu/copier"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoComment "github.com/madmuzz05/go-final-project/service/comment/dto"
	entityUser "github.com/madmuzz05/go-final-project/service/user/entity"
)

func (u *CommentUsecase) GetOne(ctx context.Context, id int) (res dtoComment.CommentResponse, err sysresponse.IError) {
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
	model, errModel := u.CommentRepository.GetOne(ctx, id)
	if errModel != nil {
		err = errModel
		return
	}
	copier.Copy(&res, &model)
	var reqUser = entityUser.User{}
	reqUser.Id = model.UserId
	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	if errUser == nil {
		res.User = &user
	}
	photo, errPhoto := u.PhotoRepository.GetOne(ctx, int(model.PhotoId))
	if errPhoto == nil {
		res.Photo = &photo
	}
	return
}

func (u *CommentUsecase) GetAll(ctx context.Context) (res []dtoComment.CommentResponse, err sysresponse.IError) {
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
	model, errModel := u.CommentRepository.GetAll(ctx)
	if errModel != nil {
		err = errModel
		return
	}

	for _, v := range model {
		var tempRes dtoComment.CommentResponse
		copier.Copy(&tempRes, &v)

		var reqUser = entityUser.User{}
		reqUser.Id = v.UserId
		user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
		if errUser == nil {
			tempRes.User = &user
		}
		photo, errPhoto := u.PhotoRepository.GetOne(ctx, int(v.PhotoId))
		if errPhoto == nil {
			tempRes.Photo = &photo
		}

		res = append(res, tempRes)
	}
	return
}

func (u *CommentUsecase) CreateComment(ctx context.Context, req dtoComment.CommentRequest) (res dtoComment.CommentResponse, err sysresponse.IError) {
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

	photo, errPhoto := u.PhotoRepository.GetOne(ctx, int(req.PhotoId))
	if errPhoto != nil {
		err = errPhoto
	}

	model, errModel := u.CommentRepository.CreateComment(ctx, req)
	if errModel != nil {
		err = errModel
		return
	}
	copier.Copy(&res, &model)
	res.User = &user
	res.Photo = &photo

	return
}
func (u *CommentUsecase) UpdateComment(ctx context.Context, req dtoComment.CommentRequest, id int) (res dtoComment.CommentResponse, err sysresponse.IError) {
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

	model, errModel := u.CommentRepository.UpdateComment(ctx, req, id)
	if errModel != nil {
		err = errModel
		return
	}

	copier.Copy(&res, &model)
	var reqUser = entityUser.User{}
	reqUser.Id = req.UserId

	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	if errUser == nil {
		res.User = &user
	}

	photo, errPhoto := u.PhotoRepository.GetOne(ctx, int(req.PhotoId))
	if errPhoto == nil {
		res.Photo = &photo
	}
	return
}

func (u *CommentUsecase) DeleteComment(ctx context.Context, id int) (err sysresponse.IError) {
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

	errModel := u.CommentRepository.DeleteComment(ctx, id)
	if errModel != nil {
		err = errModel
		return
	}
	return
}
