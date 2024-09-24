package repository

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	userDto "github.com/madmuzz05/go-final-project/service/user/dto"
	userEntity "github.com/madmuzz05/go-final-project/service/user/entity"
)

type UserRepository struct {
	gormDb *postgres.GormDB
}

func InitUserRepository(gormDb *postgres.GormDB) IUserRepository {
	return &UserRepository{
		gormDb: gormDb,
	}
}

type IUserRepository interface {
	Register(ctx context.Context, req userDto.UserRequest) (res userEntity.User, err sysresponse.IError)
	GetDataUser(ctx context.Context, req userEntity.User) (res userEntity.User, err sysresponse.IError)
}
