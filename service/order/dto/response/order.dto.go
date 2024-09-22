package response

import "time"

type OrdersDtoResponse struct {
	OrderId      uint      `json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []ItemsDtoResponse
}

type ItemsDtoResponse struct {
	ItemId      uint      `json:"id"`
	ItemCode    string    `json:"itemcode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"orderid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
