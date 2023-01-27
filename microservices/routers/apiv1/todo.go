package apiv1

import (
	"examples/microservices/controllers"
	_ "examples/microservices/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
//
//	@Summary		List todos
//	@Description	get todos
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		models.Todo
//	@Router			/api/v1/todo [get]
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
