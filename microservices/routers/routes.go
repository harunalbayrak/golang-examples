package routers

import (
	"examples/microservices/controllers"

	"github.com/gin-gonic/gin"
)

func SetupV1(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("todo", controllers.GetTodos)
		apiv1.POST("todo", controllers.CreateATodo)
		apiv1.GET("todo/:id", controllers.GetATodo)
		apiv1.PUT("todo/:id", controllers.UpdateATodo)
		apiv1.DELETE("todo/:id", controllers.DeleteATodo)
	}
}
