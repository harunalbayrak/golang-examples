package apiv1

import (
	"examples/microservices/controllers"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	controllers.GetTodos(c)
}

func CreateATodo(c *gin.Context) {
	controllers.CreateATodo(c)
}

func GetATodo(c *gin.Context) {
	controllers.GetATodo(c)
}

func UpdateATodo(c *gin.Context) {
	controllers.UpdateATodo(c)
}

func DeleteATodo(c *gin.Context) {
	controllers.DeleteATodo(c)
}
