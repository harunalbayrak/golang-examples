package controllers

import (
	"examples/microservices/models"
	"examples/microservices/pkg/e"
	"examples/microservices/pkg/setting"
	"examples/microservices/pkg/util"
	"html"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Key      string `json:"key"`
	Type     string `json:"type"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func getUserID(c *gin.Context) (userID uint) {
	return c.MustGet("id").(uint)
}

func GetUsers(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    e.ERROR_STATUS_UNAUTHORIZED,
			"message": e.GetMsg(e.ERROR_STATUS_UNAUTHORIZED),
		})
	}

	var user []models.User
	err = models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    e.ERROR_NOT_FOUND,
			"message": e.GetMsg(e.ERROR_NOT_FOUND),
		})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateAUser(c *gin.Context) {
	var user models.User
	user.Type = c.Params.ByName("type")
	if !strings.EqualFold(user.Type, "Admin") && !strings.EqualFold(user.Type, "User") {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&user)
	err := models.CreateAUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetAUser(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	id := c.Params.ByName("id")
	var user models.User
	err = models.GetAUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateAUser(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	var user models.User
	id := c.Params.ByName("id")
	err = models.GetAUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateAUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteAUser(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	var user models.User
	id := c.Params.ByName("id")
	err = models.DeleteAUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

func Register(c *gin.Context) {
	var input Auth

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = html.EscapeString(strings.TrimSpace(input.Username))
	user.Password = string(hashedPassword)
	user.Type = input.Type

	if strings.EqualFold(user.Type, "Admin") {
		if input.Key == setting.AppSettings.Key {
			RegisterUser(c, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "registiration is failed: key is wrong."})
		}
	} else if strings.EqualFold(user.Type, "User") {
		RegisterUser(c, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there is such no type"})
	}
}

func RegisterUser(c *gin.Context, user models.User) {
	err := models.CreateAUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creation is failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}

func Login(c *gin.Context) {
	var input Auth

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.CheckAuth(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	token, err := util.GenerateToken2(id)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
