package usecaseSosmed

import (
	"context"

	"github.com/jinzhu/copier"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/dto"
	entitySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/entity"
	"github.com/madmuzz05/go-final-project/service/user/dto"
)

func (u *SosmedUsecase) GetOne(ctx context.Context, id int) (res dtoSosmed.SosmedResposnse, err sysresponse.IError) {
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
	model, errModel := u.SosmedRepository.GetOne(ctx, id)
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

func (u *SosmedUsecase) GetAll(ctx context.Context) (res []dtoSosmed.SosmedResposnse, err sysresponse.IError) {
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
	model, errModel := u.SosmedRepository.GetAll(ctx)
	if errModel != nil {
		err = errModel
		return
	}

	for _, v := range model {
		var tempRes dtoSosmed.SosmedResposnse
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

func (u *SosmedUsecase) CreateSocialMedia(ctx context.Context, req entitySosmed.SosialMedia) (res dtoSosmed.SosmedResposnse, err sysresponse.IError) {
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
	model, errModel := u.SosmedRepository.CreateSocialMedia(ctx, req)
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
func (u *SosmedUsecase) UpdateSocialMedia(ctx context.Context, req entitySosmed.SosialMedia, id int) (res dtoSosmed.SosmedResposnse, err sysresponse.IError) {
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
	model, errModel := u.SosmedRepository.UpdateSocialMedia(ctx, req, id)
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

func (u *SosmedUsecase) DeleteSocialMedia(ctx context.Context, id int) (err sysresponse.IError) {
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
	errModel := u.SosmedRepository.DeleteSocialMedia(ctx, id)
	if errModel != nil {
		err = errModel
		return
	}
	return
}
