// dtoSosmed package
package dtoSosmed

// SosmedRequest schema
type SosmedRequest struct {
	Name           string `gorm:"not null" json:"name" binding:"required"`
	SosialMediaUrl string `gorm:"not null" json:"sosial_media_url" binding:"required"`
	UserId         uint   `gorm:"not null" json:"user_id" binding:"required"`
}
