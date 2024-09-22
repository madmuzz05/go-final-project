package entitySosmed

import global "github.com/madmuzz05/go-final-project/service/global/entity"

type SosialMedia struct {
	global.GormModel
	Name           string `gorm:"not null" json:"name" binding:"required"`
	SosialMediaUrl string `gorm:"not null" json:"sosial_media_url" binding:"required"`
	UserId         uint   `gorm:"not null" json:"user_id" binding:"required"`
}

func (SosialMedia) TableName() string {
	return "public.sosialmedia"
}
