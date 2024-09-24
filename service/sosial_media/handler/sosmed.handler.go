package handlerSosmed

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	validation "github.com/madmuzz05/go-final-project/pkg/validate"
	dtoSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/dto"
)

// GetOne godoc
// @Summary Get a sosmed by ID
// @Description Retrieve a single sosmed by its ID
// @Tags Sosmed
// @Accept json
// @Produce json
// @Param id path int true "Sosmed ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoSosmed.SosmedResposnse}
// @Security BearerAuth
// @Router /sosmed/getOne/{id} [get]
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

// GetAll godoc
// @Summary get all sosmed
// @Description Retrieve all sosmed
// @Tags Sosmed
// @Accept json
// @Produce json
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=[]dtoSosmed.SosmedResposnse}
// @Security BearerAuth
// @Router /sosmed/getAll [get]
func (h SosmedHandler) GetAll(ctx *gin.Context) {
	res, err := h.SosmedUsecase.GetAll(ctx)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

// CreateSosialMedia godoc
// @Summary Create sosmed
// @Description Create sosmed
// @Tags Sosmed
// @Accept json
// @Produce json
// @Param json body dtoSosmed.SosmedRequest true "body request"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoSosmed.SosmedResposnse}
// @Security BearerAuth
// @Router /sosmed/createSosialMedia [post]
func (h SosmedHandler) CreateSosialMedia(ctx *gin.Context) {
	req := dtoSosmed.SosmedRequest{}

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

// UpdateSosialMedia godoc
// @Summary Update sosmed
// @Description Update sosmed
// @Tags Sosmed
// @Accept json
// @Produce json
// @Param id path int true "Sosmed ID"
// @Param json body dtoSosmed.SosmedRequest true "body request"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoSosmed.SosmedResposnse}
// @Security BearerAuth
// @Router /sosmed/updateSosialMedia/{id} [put]
func (h SosmedHandler) UpdateSosialMedia(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	req := dtoSosmed.SosmedRequest{}
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

// DeleteSosialMedia godoc
// @Summary delete sosmed
// @Description delete sosmed
// @Tags Sosmed
// @Accept json
// @Produce json
// @Param id path int true "Sosmed ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=nil}
// @Security BearerAuth
// @Router /sosmed/deleteSosialMedia/{id} [delete]
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
