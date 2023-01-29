package routers

import (
	"examples/microservices/routers/apiv1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1.SetupAPIV1(r)

	return r
}
