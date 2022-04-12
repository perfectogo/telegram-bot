package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/perfectogo/telegram-bot/client/docs" // swag
	"github.com/perfectogo/telegram-bot/client/handlers"
	"github.com/perfectogo/telegram-bot/client/services"
	"github.com/perfectogo/telegram-bot/config"
	"github.com/perfectogo/telegram-bot/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.InterfaceServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handlers.New(&handlers.HandlerConfig{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/tg")

	api.POST("/send", handler.SendMessageHandler)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
