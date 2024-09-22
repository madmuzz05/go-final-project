package handlerComment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	validation "github.com/madmuzz05/go-final-project/pkg/validate"
	entityComment "github.com/madmuzz05/go-final-project/service/comment/entity"
)

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

func (h CommentHandler) GetAll(ctx *gin.Context) {
	res, err := h.CommentUsecase.GetAll(ctx)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

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
