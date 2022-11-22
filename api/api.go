package api

import (
	v1 "github.com/burxondv/word_task/api/v1"
	"github.com/burxondv/word_task/config"
	"github.com/burxondv/word_task/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/burxondv/word_task/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for word api
// @version         1.0
// @description     This is a word service api.
// @host      localhost:8000
// @BasePath  /v1

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	apiV1.POST("/word", handlerV1.CreateWord)
	apiV1.GET("/word", handlerV1.GetWord)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
