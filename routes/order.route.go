package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	"github.com/madmuzz05/go-final-project/service/order/handler"
	"github.com/madmuzz05/go-final-project/service/order/repository"
	"github.com/madmuzz05/go-final-project/service/order/usecase"
)

type surveyRoutes struct {
	Handler handler.OrderHandler
	Router  *gin.RouterGroup
}

func InitOrderRouter(
	router *gin.RouterGroup, gormDB *postgres.GormDB,
) *surveyRoutes {
	handler := handler.InitOrderHandler(
		usecase.InitOrderUsecase(
			repository.InitOrderRepository(
				gormDB,
			),
			gormDB,
		),
	)
	return &surveyRoutes{
		Handler: *handler,
		Router:  router,
	}
}

func (r *surveyRoutes) Routes() {
	router := r.Router.Group("/order")
	router.GET("/get-orders", r.Handler.GetOrders)
	router.POST("/store-order", r.Handler.StoreOrder)
	router.PUT("/store-order/:order_id", r.Handler.StoreOrder)
	router.DELETE("/delete-order/:order_id", r.Handler.DeleteOrder)
}
