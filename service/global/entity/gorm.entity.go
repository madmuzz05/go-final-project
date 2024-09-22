package entity

import "time"

type GormModel struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
