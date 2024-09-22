package handler

import "github.com/madmuzz05/go-final-project/service/order/usecase"

type OrderHandler struct {
	OrderUsecase usecase.IOrderUsecase
}

func InitOrderHandler(orderHandler usecase.IOrderUsecase) *OrderHandler {
	return &OrderHandler{
		OrderUsecase: orderHandler,
	}
}
