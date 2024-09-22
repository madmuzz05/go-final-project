package usecase

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/madmuzz05/go-final-project/pkg/helper/bcrypt"
	auth "github.com/madmuzz05/go-final-project/pkg/helper/jwt"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	"github.com/madmuzz05/go-final-project/service/user/dto"
	"github.com/madmuzz05/go-final-project/service/user/entity"
)

func (u *UserUsecase) Register(ctx context.Context, req dto.UserRequest) (res entity.User, err sysresponse.IError) {
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
	var emailReq = dto.UserRequest{
		Email: req.Email,
	}
	_, err = u.UserRepository.GetDataUser(ctx, emailReq)
	if err != nil && err.GetStatusCode() != http.StatusNotFound {
		return
	} else if err == nil {
		err = sysresponse.GetErrorMessage(nil, http.StatusBadRequest, "Email is registered")
		return
	}

	var usernameReq = dto.UserRequest{
		Username: req.Username,
	}
	_, err = u.UserRepository.GetDataUser(ctx, usernameReq)
	if err != nil && err.GetStatusCode() != http.StatusNotFound {
		return
	} else if err == nil {
		err = sysresponse.GetErrorMessage(nil, http.StatusBadRequest, "Username is registered")
		return
	}

	res, err = u.UserRepository.Register(ctx, req)
	if err != nil {
		return
	}
	return
}
func (u *UserUsecase) Login(ctx context.Context, req dto.LoginRequest) (res dto.LoginResponse, err sysresponse.IError) {
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
	var userReq = dto.UserRequest{
		Username: &req.Username,
	}
	user, userErr := u.UserRepository.GetDataUser(ctx, userReq)
	if userErr != nil {
		err = userErr
		return
	}

	if !bcrypt.ComparePassword([]byte(user.Password), []byte(req.Password)) {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Wrong Password")
		return
	}
	token, expiredAt := auth.GenerateToken(user.Id, user.Username)

	copier.Copy(&res, &user)
	res.Token.Token = token
	res.Token.ExpiredAt = expiredAt
	return

}

func (u *UserUsecase) GetDataUser(ctx context.Context, req dto.UserRequest) (res entity.User, err sysresponse.IError) {
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
	res, err = u.UserRepository.GetDataUser(ctx, req)
	if err != nil {
		return
	}
	return
}
