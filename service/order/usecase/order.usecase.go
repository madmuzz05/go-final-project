package usecase

import (
	"context"

	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	"github.com/madmuzz05/go-final-project/service/order/dto/request"
	"github.com/madmuzz05/go-final-project/service/order/dto/response"
)

func (u *OrderUsecase) GetOrders(ctx context.Context) (res []response.OrdersDtoResponse, err sysresponse.IError) {
	res, err = u.OrderRepository.GetOrders(ctx)
	if err != nil {
		return
	}
	return
}
func (u *OrderUsecase) StoreOrder(ctx context.Context, req request.OrderDtoRequest, id string) (res response.OrdersDtoResponse, err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	res, err = u.OrderRepository.StoreOrder(ctx, req, id)
	if err != nil {
		return
	}
	return
}

func (u *OrderUsecase) GetOrder(ctx context.Context, id string) (res response.OrdersDtoResponse, err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	res, err = u.OrderRepository.GetOrder(ctx, id)
	if err != nil {
		return
	}
	return
}

func (u *OrderUsecase) DeleteOrder(ctx context.Context, id string) (err sysresponse.IError) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()
	err = u.OrderRepository.DeleteOrder(ctx, id)
	if err != nil {
		return
	}
	return
}
