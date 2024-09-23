package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/madmuzz05/go-final-project/internal/config"
	"github.com/madmuzz05/go-final-project/internal/database/gorm/postgres"
	"github.com/madmuzz05/go-final-project/pkg/middleware"
	"github.com/madmuzz05/go-final-project/service/docs"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	HttpServer *http.Server
}

// @title MyGram API
// @version 1.0
// @description This is a api documentation for MyGram API Final Project goang Hacktive
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @host localhost:8081
// @BasePath /api/v1
func Run(ctx context.Context) error {
	gormDB, err := postgres.LoadGorm(&config.AppConfig)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("database: %s", err))
		os.Exit(1)
	}
	// Init Router
	router := InitRouter(gormDB)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s := &Server{
		HttpServer: server,
	}

	go func() {
		log.Info().Msg(fmt.Sprintf("success to listen and serve on port : %s", config.AppConfig.Port))
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve: %+v")
		}
	}()
	<-ctx.Done()

	// Gracefull Shutdown
	gracefulShutdownPeriod := 30 * time.Second
	log.Warn().Msg("shutting down http server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()
	if err := s.HttpServer.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server gracefully")
		return err
	}

	return nil
}

func SetupRouter() *gin.Engine {
	// set the runtime mode

	gin.SetMode(gin.DebugMode)

	// create a new router instance
	router := gin.Default()

	// set up middlewares
	// router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Accept", "Content-Type", "Authorization", "User-Agent"},
		ExposeHeaders:   []string{"Content-Length"},
	}))
	router.Use(gin.Recovery())

	return router
}

func InitRouter(gormDB *postgres.GormDB) *gin.Engine {
	// docs.SwaggerInfo.BasePath = "/api/v1"

	router := SetupRouter()

	api := router.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"
	api.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.URL("doc.json"),
		ginSwagger.DocExpansion("none"),
		ginSwagger.DeepLinking(true),
		ginSwagger.PersistAuthorization(true),
	))
	InitUserRouter(api, gormDB).Routes()
	api.Use(middleware.Authentication())
	InitSosmedRouter(api, gormDB).Routes()
	InitPhotoRouter(api, gormDB).Routes()
	InitCommentRouter(api, gormDB).Routes()

	return router
}
