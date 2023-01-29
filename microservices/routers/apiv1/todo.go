package apiv1

import (
	"examples/microservices/controllers"
	"examples/microservices/models"

	"github.com/gin-gonic/gin"
)

var (
	_ = models.Todo{}
)

// GetTodos godoc
//
//	@Summary		Get Todos
//	@Description	Get Todos
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Todo
//	@Router			/api/v1/todo [get]
func GetTodos(c *gin.Context) {
	controllers.GetTodos(c)
}

// CreateATodo godoc
//
//	@Summary		Create A Todo
//	@Description	Create A Todo
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	models.Todo
//	@Router			/api/v1/todo [post]
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
