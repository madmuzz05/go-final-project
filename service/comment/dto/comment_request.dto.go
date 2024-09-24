// dtoComment package
package dtoComment

// CommentRequest schema
type CommentRequest struct {
	Message string `gorm:"not null" json:"message" binding:"required"`
	PhotoId uint   `gorm:"not null" json:"photo_id" binding:"required"`
	UserId  uint   `gorm:"not null" json:"user_id" binding:"required"`
}
