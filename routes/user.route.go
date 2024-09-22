package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	"github.com/madmuzz05/go-final-project/service/user/handler"
	"github.com/madmuzz05/go-final-project/service/user/repository"
	"github.com/madmuzz05/go-final-project/service/user/usecase"
)

type userRoutes struct {
	Handler handler.UserHandler
	Router  *gin.RouterGroup
}

func InitUserRouter(
	router *gin.RouterGroup, gormDB *postgres.GormDB,
) *userRoutes {
	handler := handler.InitUserHandler(
		usecase.InitUserUsecase(
			repository.InitUserRepository(
				gormDB,
			),
			gormDB,
		),
	)
	return &userRoutes{
		Handler: *handler,
		Router:  router,
	}
}

func (r *userRoutes) Routes() {
	router := r.Router.Group("/user")
	router.POST("/register", r.Handler.Register)
	router.POST("/login", r.Handler.Login)
}
