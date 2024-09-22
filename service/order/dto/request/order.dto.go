package request

type OrderDtoRequest struct {
	OrderedAt    string           `json:"orderedAt"`
	CustomerName string           `json:"customerName" binding:"required"`
	Items        []ItemDtoRequest `json:"items"`
}

type ItemDtoRequest struct {
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
