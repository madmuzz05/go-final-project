// entityComment package
package entityComment

import entityGlobal "github.com/madmuzz05/go-final-project/service/global/entity"

// Comment schema
type Comment struct {
	entityGlobal.GormModel
	Message string `gorm:"not null" json:"message" binding:"required"`
	PhotoId uint   `gorm:"not null" json:"photo_id" binding:"required"`
	UserId  uint   `gorm:"not null" json:"user_id" binding:"required"`
}

func (Comment) TableName() string {
	return "public.comment"
}
