package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	"github.com/madmuzz05/go-final-project/pkg/middleware"
	handlerSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/handler"
	repositorySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/repository"
	usecaseSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/usecase"
	repositoryUser "github.com/madmuzz05/go-final-project/service/user/repository"
)

type sosmedRoutes struct {
	Handler handlerSosmed.SosmedHandler
	Router  *gin.RouterGroup
	GormDB  *postgres.GormDB
}

func InitSosmedRouter(
	router *gin.RouterGroup, gormDB *postgres.GormDB,
) *sosmedRoutes {
	handler := handlerSosmed.InitSosmedHandler(
		usecaseSosmed.InitSosmedUsecase(
			repositorySosmed.InitSosmedRepository(
				gormDB,
			),
			repositoryUser.InitUserRepository(
				gormDB,
			),
			gormDB,
		),
	)
	db := gormDB
	return &sosmedRoutes{
		Handler: *handler,
		Router:  router,
		GormDB:  db,
	}
}

func (r *sosmedRoutes) Routes() {
	tableName := "public.sosialmedia"
	router := r.Router.Group("/sosmed")
	router.GET("/getOne/:id", r.Handler.GetOne)
	router.GET("/getAll", r.Handler.GetAll)
	router.POST("/createSosialMedia", r.Handler.CreateSosialMedia)
	router.PUT("/updateSosialMedia/:id", middleware.InitAuthorizationMiddleware(r.GormDB).Authorization("id", tableName), r.Handler.UpdateSosialMedia)
	router.DELETE("/deleteSosialMedia/:id", middleware.InitAuthorizationMiddleware(r.GormDB).Authorization("id", tableName), r.Handler.DeleteSosialMedia)
}
