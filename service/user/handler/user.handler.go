package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	validation "github.com/madmuzz05/go-final-project/pkg/validate"

	"github.com/madmuzz05/go-final-project/service/user/dto"
)

func (h UserHandler) Register(ctx *gin.Context) {
	// req := dto.GenerateQrCodeRequest{}
	request := dto.UserRequest{}

	if valErr := validation.ValidateRequest(ctx, &request); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.UserUsecase.Register(ctx, request)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusCreated, "Register Success", res)
}

func (h UserHandler) Login(ctx *gin.Context) {
	// req := dto.GenerateQrCodeRequest{}
	request := dto.LoginRequest{}

	if valErr := validation.ValidateRequest(ctx, &request); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, valErr.GetMessage(), nil)
		return
	}

	res, err := h.UserUsecase.Login(ctx, request)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Login Success", res)
}
