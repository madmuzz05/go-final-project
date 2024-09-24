package handlerPhoto

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	validation "github.com/madmuzz05/go-final-project/pkg/validate"
	dtoPhoto "github.com/madmuzz05/go-final-project/service/photo/dto"
)

// GetOne godoc
// @Summary Get a photo by ID
// @Description Retrieve a single photo by its ID
// @Tags Photo
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoPhoto.PhotoResponse}
// @Security BearerAuth
// @Router /photo/getOne/{id} [get]
func (h PhotoHandler) GetOne(ctx *gin.Context) {
	// req := dto.GenerateQrCodeRequest{}
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	res, err := h.PhotoUsecase.GetOne(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

// GetAll godoc
// @Summary get all photo
// @Description Retrieve all photo
// @Tags Photo
// @Accept json
// @Produce json
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=[]dtoPhoto.PhotoResponse}
// @Security BearerAuth
// @Router /photo/getAll [get]
func (h PhotoHandler) GetAll(ctx *gin.Context) {
	res, err := h.PhotoUsecase.GetAll(ctx)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data.", res)
}

// CreatePhoto godoc
// @Summary Create photo
// @Description Create photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param json body dtoPhoto.PhotoRequest true "body request"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoPhoto.PhotoResponse}
// @Security BearerAuth
// @Router /photo/createPhoto [post]
func (h PhotoHandler) CreatePhoto(ctx *gin.Context) {
	req := dtoPhoto.PhotoRequest{}

	if valErr := validation.ValidateRequest(ctx, &req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.PhotoUsecase.CreatePhoto(ctx, req)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusCreated, "Create Data Success.", res)
}

// UpdatePhoto godoc
// @Summary update photo
// @Description update photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param json body dtoPhoto.PhotoRequest true "body request"
// @Param id path int true "Photo ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=dtoPhoto.PhotoResponse}
// @Security BearerAuth
// @Router /photo/updatePhoto/{id} [put]
func (h PhotoHandler) UpdatePhoto(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	req := dtoPhoto.PhotoRequest{}
	if valErr := validation.ValidateRequest(ctx, &req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.PhotoUsecase.UpdatePhoto(ctx, req, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Update Data Success.", res)
}

// DeletePhoto godoc
// @Summary delete photo
// @Description delete photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param json body dtoPhoto.PhotoRequest true "body request"
// @Param id path int true "Photo ID"
// @Success 200 {object} sysresponse.Success{status=int,success=bool,message=string,data=nil}
// @Security BearerAuth
// @Router /photo/deletePhoto/{id} [delete]
func (h PhotoHandler) DeletePhoto(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Invalid parameter", gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	err := h.PhotoUsecase.DeletePhoto(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Delete Data Succcess.", nil)
}
