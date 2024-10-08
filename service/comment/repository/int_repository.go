package repositoryComment

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoComment "github.com/madmuzz05/go-final-project/service/comment/dto"
	entityComment "github.com/madmuzz05/go-final-project/service/comment/entity"
)

type CommentRepository struct {
	gormDb *postgres.GormDB
}

func InitCommentRepository(gormDb *postgres.GormDB) ICommentRepository {
	return &CommentRepository{
		gormDb: gormDb,
	}
}

type ICommentRepository interface {
	GetAll(ctx context.Context) (res []entityComment.Comment, err sysresponse.IError)
	GetOne(ctx context.Context, id int) (res entityComment.Comment, err sysresponse.IError)
	GetCommentByPhotoId(ctx context.Context, id int) (res []dtoComment.CommentPhotoResponse, err sysresponse.IError)
	CreateComment(ctx context.Context, req dtoComment.CommentRequest) (res entityComment.Comment, err sysresponse.IError)
	UpdateComment(ctx context.Context, req dtoComment.CommentRequest, id int) (res entityComment.Comment, err sysresponse.IError)
	DeleteComment(ctx context.Context, id int) (err sysresponse.IError)
}
