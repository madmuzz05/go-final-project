package handlerComment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	validation "github.com/madmuzz05/go-final-project/pkg/validate"
	entityComment "github.com/madmuzz05/go-final-project/service/comment/entity"
)

// GetOne godoc
// @Summary Get a comment by ID
// @Description Retrieve a single comment by its ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoComment.CommentResponse}
// @Security BearerAuth
// @Router /comment/getOne/{id} [get]
func (h CommentHandler) GetOne(ctx *gin.Context) {
	// req := dto.GenerateQrCodeRequest{}
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	res, err := h.CommentUsecase.GetOne(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

// GetAll godoc
// @Summary get all comment
// @Description Retrieve all comment
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=[]dtoComment.CommentResponse}
// @Security BearerAuth
// @Router /comment/getAll [get]
func (h CommentHandler) GetAll(ctx *gin.Context) {
	res, err := h.CommentUsecase.GetAll(ctx)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

// CreateComment godoc
// @Summary Create comment
// @Description Create comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param json body entityComment.Comment true "body request"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoComment.CommentResponse}
// @Security BearerAuth
// @Router /comment/createComment [post]
func (h CommentHandler) CreateComment(ctx *gin.Context) {
	req := entityComment.Comment{}

	if valErr := validation.ValidateRequest(ctx, &req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.CommentUsecase.CreateComment(ctx, req)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusCreated, "Create Data Success.", res)
}

// UpdateComment godoc
// @Summary update comment
// @Description update comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param json body entityComment.Comment true "body request"
// @Param id path int true "Comment ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoComment.CommentResponse}
// @Security BearerAuth
// @Router /comment/updateComment/{id} [put]
func (h CommentHandler) UpdateComment(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	req := entityComment.Comment{}
	if valErr := validation.ValidateRequest(ctx, &req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.CommentUsecase.UpdateComment(ctx, req, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Update Data Success.", res)
}

// DeleteComment godoc
// @Summary delete comment
// @Description delete comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=nil}
// @Security BearerAuth
// @Router /comment/deleteComment/{id} [delete]
func (h CommentHandler) DeleteComment(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	err := h.CommentUsecase.DeleteComment(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Delete Data Succcess.", nil)
}
