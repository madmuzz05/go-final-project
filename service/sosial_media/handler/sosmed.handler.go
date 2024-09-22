package handlerSosmed

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	validation "github.com/madmuzz05/go-final-project/pkg/validate"
	entitySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/entity"
)

func (h SosmedHandler) GetOne(ctx *gin.Context) {
	// req := dto.GenerateQrCodeRequest{}
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	res, err := h.SosmedUsecase.GetOne(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

func (h SosmedHandler) GetAll(ctx *gin.Context) {
	res, err := h.SosmedUsecase.GetAll(ctx)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

func (h SosmedHandler) CreateSosialMedia(ctx *gin.Context) {
	req := entitySosmed.SosialMedia{}

	if valErr := validation.ValidateRequest(ctx, &req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.SosmedUsecase.CreateSocialMedia(ctx, req)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusCreated, "Create Data Success.", res)
}

func (h SosmedHandler) UpdateSosialMedia(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	req := entitySosmed.SosialMedia{}
	if valErr := validation.ValidateRequest(ctx, &req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.SosmedUsecase.UpdateSocialMedia(ctx, req, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Update Data Success.", res)
}

func (h SosmedHandler) DeleteSosialMedia(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	err := h.SosmedUsecase.DeleteSocialMedia(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Delete Data Succcess.", nil)
}
