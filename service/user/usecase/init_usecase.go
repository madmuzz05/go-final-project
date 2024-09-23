package usecase

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"

	"github.com/madmuzz05/go-final-project/service/user/dto"
	entityUser "github.com/madmuzz05/go-final-project/service/user/entity"
	userRepo "github.com/madmuzz05/go-final-project/service/user/repository"
)

type UserUsecase struct {
	UserRepository userRepo.IUserRepository
	GormDB         *postgres.GormDB
}

func InitUserUsecase(userRepository userRepo.IUserRepository, gormDb *postgres.GormDB) IUserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
		GormDB:         gormDb,
	}
}

type IUserUsecase interface {
	Register(ctx context.Context, req dto.UserRequest) (res entityUser.User, err sysresponse.IError)
	Login(ctx context.Context, req dto.LoginRequest) (res dto.LoginResponse, err sysresponse.IError)
	GetDataUser(ctx context.Context, req dto.UserRequest) (res entityUser.User, err sysresponse.IError)
}
