package entity

import (
	"time"
)

type Order struct {
	OrderId      uint      `json:"order_id" gorm:"primary_key"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Item    `gorm:"foreignKey:OrderId;references:OrderId"`
}

func (Order) TableName() string {
	return "assignment2.order"
}

type Item struct {
	ItemId      uint      `json:"item_id" gorm:"primary_key"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description" gorm:"type:text"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OrderId     uint
}

func (Item) TableName() string {
	return "assignment2.item"
}
