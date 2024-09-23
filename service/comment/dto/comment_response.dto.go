// dtoComment package
package dtoComment

import (
	entityGlobal "github.com/madmuzz05/go-final-project/service/global/entity"
	entityPhoto "github.com/madmuzz05/go-final-project/service/photo/entity"
	entityUser "github.com/madmuzz05/go-final-project/service/user/entity"
)

// CommentResponse schema
type CommentResponse struct {
	entityGlobal.GormModel
	Message string             `gorm:"not null" json:"message"`
	PhotoId uint               `gorm:"not null" json:"photo_id"`
	UserId  uint               `gorm:"not null" json:"user_id"`
	User    *entityUser.User   `json:"user"`
	Photo   *entityPhoto.Photo `json:"photo"`
}

// CommentPhotoResponse schema
type CommentPhotoResponse struct {
	Message  string  `gorm:"not null" json:"message"`
	PhotoId  uint    `gorm:"not null" json:"photo_id"`
	UserId   uint    `gorm:"not null" json:"user_id"`
	Username *string `json:"username"`
}
