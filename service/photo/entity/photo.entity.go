package entityPhoto

import global "github.com/madmuzz05/go-final-project/service/global/entity"

type Photo struct {
	global.GormModel
	Title    string `gorm:"not null" json:"title" binding:"required"`
	Caption  string `gorm:"not null" json:"caption" binding:"required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" binding:"required"`
	UserId   uint   `gorm:"not null" json:"user_id" binding:"required"`
}

func (Photo) TableName() string {
	return "public.photo"
}
