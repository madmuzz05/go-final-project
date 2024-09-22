package entity

import (
	global "github.com/madmuzz05/go-final-project/service/global/entity"
)

type User struct {
	global.GormModel
	Username string `gorm:"not null;uniqueindex" json:"username" binding:"required"`
	Email    string `gorm:"not null;uniqueindex" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=6"`
	Age      int    `gorm:"not null" json:"age" binding:"required,gte=8"`
}

func (User) TableName() string {
	return "public.user"
}