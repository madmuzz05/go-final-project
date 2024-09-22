package usecasePhoto

import (
	"context"

	"github.com/jinzhu/copier"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoPhoto "github.com/madmuzz05/go-final-project/service/photo/dto"
	entityPhoto "github.com/madmuzz05/go-final-project/service/photo/entity"
	"github.com/madmuzz05/go-final-project/service/user/dto"
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
	var reqUser = dto.UserRequest{
		Id: int(model.UserId),
	}
	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	copier.Copy(&res, &model)
	if errUser == nil {
		res.User = &user
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

		var reqUser = dto.UserRequest{
			Id: int(v.UserId),
		}
		user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
		if errUser == nil {
			tempRes.User = &user
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
	model, errModel := u.PhotoRepository.CreatePhoto(ctx, req)
	if errModel != nil {
		err = errModel
		return
	}
	var reqUser = dto.UserRequest{
		Id: int(model.UserId),
	}
	user, errUser := u.UserRepository.GetDataUser(ctx, reqUser)
	copier.Copy(&res, &model)
	if errUser == nil {
		res.User = &user
	}
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
	var reqUser = dto.UserRequest{
		Id: int(model.UserId),
	}
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
