// dtoSosmed package
package dtoSosmed

import (
	entityGlobal "github.com/madmuzz05/go-final-project/service/global/entity"
	entityUser "github.com/madmuzz05/go-final-project/service/user/entity"
)

// SosmedResposnse schema
type SosmedResposnse struct {
	entityGlobal.GormModel
	Name           string           `gorm:"not null" json:"name"`
	SosialMediaUrl string           `gorm:"not null" json:"sosial_media_url"`
	UserId         uint             `gorm:"not null" json:"user_id"`
	User           *entityUser.User `json:"user"`
}
