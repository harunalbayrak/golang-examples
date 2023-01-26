package routers

import (
	"examples/microservices/routers/api"

	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	SetupV1(r)

	return r
}
