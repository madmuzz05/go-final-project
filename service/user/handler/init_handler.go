package handler

import "github.com/madmuzz05/go-final-project/service/user/usecase"

type UserHandler struct {
	UserUsecase usecase.IUserUsecase
}

func InitUserHandler(userUsecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}
