// entityPhoto package
package entityPhoto

import entityGlobal "github.com/madmuzz05/go-final-project/service/global/entity"

// Photo schema
type Photo struct {
	entityGlobal.GormModel
	Title    string `gorm:"not null" json:"title" binding:"required"`
	Caption  string `gorm:"not null" json:"caption" binding:"required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" binding:"required"`
	UserId   uint   `gorm:"not null" json:"user_id" binding:"required"`
}

func (Photo) TableName() string {
	return "public.photo"
}
