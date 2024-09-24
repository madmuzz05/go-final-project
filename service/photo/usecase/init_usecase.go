package usecasePhoto

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	repositoryComment "github.com/madmuzz05/go-final-project/service/comment/repository"
	dtoPhoto "github.com/madmuzz05/go-final-project/service/photo/dto"
	repositoryPhoto "github.com/madmuzz05/go-final-project/service/photo/repository"
	repositoryUser "github.com/madmuzz05/go-final-project/service/user/repository"
)

type PhotoUsecase struct {
	PhotoRepository   repositoryPhoto.IPhotoRepository
	UserRepository    repositoryUser.IUserRepository
	CommentRepository repositoryComment.ICommentRepository
	GormDB            *postgres.GormDB
}

func InitPhotoUsecase(photoRepository repositoryPhoto.IPhotoRepository, userRepository repositoryUser.IUserRepository, commentRepository repositoryComment.ICommentRepository, gormDb *postgres.GormDB) IPhotoUsecase {
	return &PhotoUsecase{
		PhotoRepository:   photoRepository,
		UserRepository:    userRepository,
		CommentRepository: commentRepository,
		GormDB:            gormDb,
	}
}

type IPhotoUsecase interface {
	GetOne(ctx context.Context, id int) (res dtoPhoto.PhotoResponse, err sysresponse.IError)
	GetAll(ctx context.Context) (res []dtoPhoto.PhotoResponse, err sysresponse.IError)
	CreatePhoto(ctx context.Context, req dtoPhoto.PhotoRequest) (res dtoPhoto.PhotoResponse, err sysresponse.IError)
	UpdatePhoto(ctx context.Context, req dtoPhoto.PhotoRequest, id int) (res dtoPhoto.PhotoResponse, err sysresponse.IError)
	DeletePhoto(ctx context.Context, id int) (err sysresponse.IError)
}
