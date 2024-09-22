package handlerComment

import (
	usecaseComment "github.com/madmuzz05/go-final-project/service/Comment/usecase"
)

type CommentHandler struct {
	CommentUsecase usecaseComment.ICommentUsecase
}

func InitCommentHandler(commentUsecase usecaseComment.ICommentUsecase) *CommentHandler {
	return &CommentHandler{
		CommentUsecase: commentUsecase,
	}
}