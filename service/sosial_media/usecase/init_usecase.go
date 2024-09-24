package usecaseSosmed

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/dto"
	repositorySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/repository"
	repositoryUser "github.com/madmuzz05/go-final-project/service/user/repository"
)

type SosmedUsecase struct {
	SosmedRepository repositorySosmed.ISosmedRepository
	UserRepository   repositoryUser.IUserRepository
	GormDB           *postgres.GormDB
}

func InitSosmedUsecase(sosmedRepository repositorySosmed.ISosmedRepository, userRepository repositoryUser.IUserRepository, gormDb *postgres.GormDB) ISosmedUsecase {
	return &SosmedUsecase{
		SosmedRepository: sosmedRepository,
		UserRepository:   userRepository,
		GormDB:           gormDb,
	}
}

type ISosmedUsecase interface {
	GetOne(ctx context.Context, id int) (res dtoSosmed.SosmedResposnse, err sysresponse.IError)
	GetAll(ctx context.Context) (res []dtoSosmed.SosmedResposnse, err sysresponse.IError)
	CreateSocialMedia(ctx context.Context, req dtoSosmed.SosmedRequest) (res dtoSosmed.SosmedResposnse, err sysresponse.IError)
	UpdateSocialMedia(ctx context.Context, req dtoSosmed.SosmedRequest, id int) (res dtoSosmed.SosmedResposnse, err sysresponse.IError)
	DeleteSocialMedia(ctx context.Context, id int) (err sysresponse.IError)
}
