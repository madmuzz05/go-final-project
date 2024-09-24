// dtoPhoto package
package dtoPhoto

// PhotoRequest schema
type PhotoRequest struct {
	Title    string `gorm:"not null" json:"title" binding:"required"`
	Caption  string `gorm:"not null" json:"caption" binding:"required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" binding:"required"`
	UserId   uint   `gorm:"not null" json:"user_id" binding:"required"`
}
