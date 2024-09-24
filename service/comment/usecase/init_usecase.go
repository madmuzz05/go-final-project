package usecaseComment

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoComment "github.com/madmuzz05/go-final-project/service/comment/dto"
	repositoryComment "github.com/madmuzz05/go-final-project/service/comment/repository"
	repositoryPhoto "github.com/madmuzz05/go-final-project/service/photo/repository"
	repositoryUser "github.com/madmuzz05/go-final-project/service/user/repository"
)

type CommentUsecase struct {
	CommentRepository repositoryComment.ICommentRepository
	UserRepository    repositoryUser.IUserRepository
	PhotoRepository   repositoryPhoto.IPhotoRepository
	GormDB            *postgres.GormDB
}

func InitCommentUsecase(
	commentRepository repositoryComment.ICommentRepository,
	userRepository repositoryUser.IUserRepository,
	photoRepository repositoryPhoto.IPhotoRepository,
	gormDb *postgres.GormDB) ICommentUsecase {
	return &CommentUsecase{
		CommentRepository: commentRepository,
		UserRepository:    userRepository,
		PhotoRepository:   photoRepository,
		GormDB:            gormDb,
	}
}

type ICommentUsecase interface {
	GetOne(ctx context.Context, id int) (res dtoComment.CommentResponse, err sysresponse.IError)
	GetAll(ctx context.Context) (res []dtoComment.CommentResponse, err sysresponse.IError)
	CreateComment(ctx context.Context, req dtoComment.CommentRequest) (res dtoComment.CommentResponse, err sysresponse.IError)
	UpdateComment(ctx context.Context, req dtoComment.CommentRequest, id int) (res dtoComment.CommentResponse, err sysresponse.IError)
	DeleteComment(ctx context.Context, id int) (err sysresponse.IError)
}
