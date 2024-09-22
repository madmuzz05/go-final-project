package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	"github.com/madmuzz05/go-final-project/pkg/middleware"
	handlerPhoto "github.com/madmuzz05/go-final-project/service/photo/handler"
	repositoryPhoto "github.com/madmuzz05/go-final-project/service/photo/repository"
	usecasePhoto "github.com/madmuzz05/go-final-project/service/photo/usecase"
	repositoryUser "github.com/madmuzz05/go-final-project/service/user/repository"
)

type photoRoutes struct {
	Handler handlerPhoto.PhotoHandler
	Router  *gin.RouterGroup
	GormDB  *postgres.GormDB
}

func InitPhotoRouter(
	router *gin.RouterGroup, gormDB *postgres.GormDB,
) *photoRoutes {
	handler := handlerPhoto.InitPhotoHandler(
		usecasePhoto.InitPhotoUsecase(
			repositoryPhoto.InitPhotoRepository(
				gormDB,
			),
			repositoryUser.InitUserRepository(
				gormDB,
			),
			gormDB,
		),
	)
	db := gormDB
	return &photoRoutes{
		Handler: *handler,
		Router:  router,
		GormDB:  db,
	}
}

func (r *photoRoutes) Routes() {
	tableName := "public.photo"
	router := r.Router.Group("/photo")
	router.GET("/getOne/:id", r.Handler.GetOne)
	router.GET("/getAll", r.Handler.GetAll)
	router.POST("/createPhoto", r.Handler.CreatePhoto)
	router.PUT("/updatePhoto/:id", middleware.InitAuthorizationMiddleware(r.GormDB).Authorization("id", tableName), r.Handler.UpdatePhoto)
	router.DELETE("/deletePhoto/:id", middleware.InitAuthorizationMiddleware(r.GormDB).Authorization("id", tableName), r.Handler.DeletePhoto)
}
