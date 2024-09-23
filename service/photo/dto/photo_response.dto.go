// dtoPhoto package
package dtoPhoto

import (
	dtoComment "github.com/madmuzz05/go-final-project/service/comment/dto"
	entityGlobal "github.com/madmuzz05/go-final-project/service/global/entity"
	entityUser "github.com/madmuzz05/go-final-project/service/user/entity"
)

// PhotoResponse schema
type PhotoResponse struct {
	entityGlobal.GormModel
	Title    string                             `gorm:"not null" json:"title"`
	Caption  string                             `gorm:"not null" json:"caption"`
	PhotoUrl string                             `gorm:"not null" json:"photo_url"`
	UserId   uint                               `gorm:"not null" json:"user_id"`
	User     *entityUser.User                   `json:"user"`
	Comment  *[]dtoComment.CommentPhotoResponse `json:"comments"`
}
