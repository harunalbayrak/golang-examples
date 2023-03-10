package apiv1

import (
	middleware "examples/microservices/middleware/jwt"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupAPIV1(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public := r.Group("/api/v1")
	{
		public.POST("/register", Register)
		public.POST("/login", Login)
		public.GET("/user", CurrentUser)
	}

	users := r.Group("/api/v1/users")
	users.Use(middleware.JwtAuthMiddleware())
	{
		users.GET("/:id", GetAUser)
		users.GET("/:id/todo", GetTodos)
		users.POST("/:id/todo", CreateATodo)
		users.GET("/:id/todo/:todo_id", GetATodo)
		users.PUT("/:id/todo/:todo_id", UpdateATodo)
		users.DELETE("/:id/todo/:todo_id", DeleteATodo)
	}

	protected := r.Group("/api/v1/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	{
		protected.GET("/users", GetUsers)
		protected.GET("/users/:id", GetAUser)
	}
}
