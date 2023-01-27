package apiv1

import (
	"examples/microservices/controllers"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	controllers.GetUsers(c)
}

func CreateAUser(c *gin.Context) {
	controllers.CreateAUser(c)
}

func GetAUser(c *gin.Context) {
	controllers.GetAUser(c)
}

func UpdateAUser(c *gin.Context) {
	controllers.UpdateAUser(c)
}

func DeleteAUser(c *gin.Context) {
	controllers.DeleteAUser(c)
}

func Register(c *gin.Context) {
	controllers.Register(c)
}

func Login(c *gin.Context) {
	controllers.Login(c)
}

func CurrentUser(c *gin.Context) {
	controllers.CurrentUser(c)
}
