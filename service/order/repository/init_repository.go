package repository

import (
	"context"

	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	"github.com/madmuzz05/go-final-project/service/order/dto/request"
	"github.com/madmuzz05/go-final-project/service/order/dto/response"
)

type OrderRepository struct {
	gormDb *postgres.GormDB
}

func InitOrderRepository(gormDb *postgres.GormDB) IOrderRepository {
	return &OrderRepository{
		gormDb: gormDb,
	}
}

type IOrderRepository interface {
	GetOrders(ctx context.Context) (res []response.OrdersDtoResponse, err sysresponse.IError)
	GetOrder(ctx context.Context, id string) (res response.OrdersDtoResponse, err sysresponse.IError)
	DeleteOrder(ctx context.Context, id string) (err sysresponse.IError)
	StoreOrder(ctx context.Context, req request.OrderDtoRequest, id string) (res response.OrdersDtoResponse, err sysresponse.IError)
}
