package repositorySosmed

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/dto"
	entitySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/entity"
)

type SosmedRepository struct {
	gormDb *postgres.GormDB
}

func InitSosmedRepository(gormDb *postgres.GormDB) ISosmedRepository {
	return &SosmedRepository{
		gormDb: gormDb,
	}
}

type ISosmedRepository interface {
	GetAll(ctx context.Context) (res []entitySosmed.SosialMedia, err sysresponse.IError)
	GetOne(ctx context.Context, id int) (res entitySosmed.SosialMedia, err sysresponse.IError)
	CreateSocialMedia(ctx context.Context, req dtoSosmed.SosmedRequest) (res entitySosmed.SosialMedia, err sysresponse.IError)
	UpdateSocialMedia(ctx context.Context, req dtoSosmed.SosmedRequest, id int) (res entitySosmed.SosialMedia, err sysresponse.IError)
	DeleteSocialMedia(ctx context.Context, id int) (err sysresponse.IError)
}
