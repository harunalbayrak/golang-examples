package apiv1

import (
	middleware "examples/microservices/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func SetupAPIV1(r *gin.Engine) {
	public := r.Group("/api/v1")
	{
		public.POST("/register", Register)
		public.POST("/login", Login)
	}

	protected := r.Group("/api/v1/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	{
		protected.GET("/user", CurrentUser)
		protected.GET("/users", GetUsers)
		protected.GET("/user/:id", GetAUser)
		protected.GET("/todo", GetTodos)
		protected.POST("/todo", CreateATodo)
		protected.GET("/todo/:id", GetATodo)
		protected.PUT("/todo/:id", UpdateATodo)
		protected.DELETE("/todo/:id", DeleteATodo)
	}
}
