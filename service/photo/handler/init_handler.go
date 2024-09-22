package handlerPhoto

import (
	usecasePhoto "github.com/madmuzz05/go-final-project/service/photo/usecase"
)

type PhotoHandler struct {
	PhotoUsecase usecasePhoto.IPhotoUsecase
}

func InitPhotoHandler(photoUsecase usecasePhoto.IPhotoUsecase) *PhotoHandler {
	return &PhotoHandler{
		PhotoUsecase: photoUsecase,
	}
}
