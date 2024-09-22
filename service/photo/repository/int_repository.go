package repositoryPhoto

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	entityPhoto "github.com/madmuzz05/go-final-project/service/photo/entity"
)

type PhotoRepository struct {
	gormDb *postgres.GormDB
}

func InitPhotoRepository(gormDb *postgres.GormDB) IPhotoRepository {
	return &PhotoRepository{
		gormDb: gormDb,
	}
}

type IPhotoRepository interface {
	GetAll(ctx context.Context) (res []entityPhoto.Photo, err sysresponse.IError)
	GetOne(ctx context.Context, id int) (res entityPhoto.Photo, err sysresponse.IError)
	CreatePhoto(ctx context.Context, req entityPhoto.Photo) (res entityPhoto.Photo, err sysresponse.IError)
	UpdatePhoto(ctx context.Context, req entityPhoto.Photo, id int) (res entityPhoto.Photo, err sysresponse.IError)
	DeletePhoto(ctx context.Context, id int) (err sysresponse.IError)
}
