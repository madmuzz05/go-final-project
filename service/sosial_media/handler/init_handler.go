package handlerSosmed

import (
	usecaseSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/usecase"
)

type SosmedHandler struct {
	SosmedUsecase usecaseSosmed.ISosmedUsecase
}

func InitSosmedHandler(sosmedUsecase usecaseSosmed.ISosmedUsecase) *SosmedHandler {
	return &SosmedHandler{
		SosmedUsecase: sosmedUsecase,
	}
}
