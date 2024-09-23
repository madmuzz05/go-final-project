package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	"github.com/madmuzz05/go-final-project/pkg/middleware"
	usecaseComment "github.com/madmuzz05/go-final-project/service/comment/usecase"
	handlerComment "github.com/madmuzz05/go-final-project/service/comment/handler"
	repositoryComment "github.com/madmuzz05/go-final-project/service/comment/repository"
	repositoryPhoto "github.com/madmuzz05/go-final-project/service/photo/repository"
	repositoryUser "github.com/madmuzz05/go-final-project/service/user/repository"
)

type commentRoutes struct {
	Handler handlerComment.CommentHandler
	Router  *gin.RouterGroup
	GormDB  *postgres.GormDB
}

func InitCommentRouter(
	router *gin.RouterGroup, gormDB *postgres.GormDB,
) *commentRoutes {
	handler := handlerComment.InitCommentHandler(
		usecaseComment.InitCommentUsecase(
			repositoryComment.InitCommentRepository(
				gormDB,
			),
			repositoryUser.InitUserRepository(
				gormDB,
			),
			repositoryPhoto.InitPhotoRepository(
				gormDB,
			),
			gormDB,
		),
	)
	db := gormDB
	return &commentRoutes{
		Handler: *handler,
		Router:  router,
		GormDB:  db,
	}
}

func (r *commentRoutes) Routes() {
	tableName := "public.comment"
	router := r.Router.Group("/comment")
	router.GET("/getOne/:id", r.Handler.GetOne)
	router.GET("/getAll", r.Handler.GetAll)
	router.POST("/createComment", r.Handler.CreateComment)
	router.PUT("/updateComment/:id", middleware.InitAuthorizationMiddleware(r.GormDB).Authorization("id", tableName), r.Handler.UpdateComment)
	router.DELETE("/deleteComment/:id", middleware.InitAuthorizationMiddleware(r.GormDB).Authorization("id", tableName), r.Handler.DeleteComment)
}
