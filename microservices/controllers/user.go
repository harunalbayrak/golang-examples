package controllers

import (
	"examples/microservices/models"
	"examples/microservices/pkg/app"
	"examples/microservices/pkg/e"
	"examples/microservices/pkg/setting"
	"examples/microservices/pkg/util"
	"fmt"
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
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	var user []models.User

	err = models.GetAllUsers(&user)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_GET_USERS)
	} else {
		app.ResponseSuccess(c, user)
	}
}

func CreateAUser(c *gin.Context) {
	var user models.User
	user.Type = c.Params.ByName("type")
	if !strings.EqualFold(user.Type, "Admin") && !strings.EqualFold(user.Type, "User") {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_INVALID_TYPE)
		return
	}

	c.BindJSON(&user)

	err := models.CreateAUser(&user)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_CREATE_USER)
	} else {
		app.ResponseSuccess(c, user)
	}
}

func GetAUser(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	id := c.Params.ByName("id")
	var user models.User

	err = models.GetAUser(&user, id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_GET_USER)
	} else {
		app.ResponseSuccess(c, user)
	}
}

func UpdateAUser(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	var user models.User
	id := c.Params.ByName("id")
	err = models.GetAUser(&user, id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_GET_USER)
	}
	c.BindJSON(&user)

	err = models.UpdateAUser(&user, id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_UPDATE_USER)
	} else {
		app.ResponseSuccess(c, user)
	}
}

func DeleteAUser(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	var user models.User
	id := c.Params.ByName("id")

	err = models.DeleteAUser(&user, id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_DELETE_USER)
	} else {
		app.ResponseSuccess(c, fmt.Sprintf("Deleted user: %s", id))
	}
}

func Register(c *gin.Context) {
	var input Auth

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("asdasd")
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_INVALID_PARAMS)
		return
	}

	user := models.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_GENERATE_HASH)
		return
	}

	user.Username = html.EscapeString(strings.TrimSpace(input.Username))
	user.Password = string(hashedPassword)
	user.Type = input.Type

	if strings.EqualFold(user.Type, "Admin") {
		if input.Key == setting.AppSettings.GeneralSettings.Key {
			RegisterUser(c, user)
		} else {
			app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_INVALID_KEY)
		}
	} else if strings.EqualFold(user.Type, "User") {
		RegisterUser(c, user)
	} else {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_INVALID_TYPE)
	}
}

func RegisterUser(c *gin.Context, user models.User) {
	err := models.CreateAUser(&user)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_REGISTER_USER)
		return
	}

	app.ResponseSuccess(c, "Registration success!")
}

func Login(c *gin.Context) {
	var input Auth

	if err := c.ShouldBindJSON(&input); err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_INVALID_PARAMS)
		return
	}

	id, err := models.CheckAuth(input.Username, input.Password)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_WRONG_USERNAME_OR_PASSWORD)
		return
	}

	token, err := util.GenerateToken(id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_TOKEN_GENERATION_FAIL)
		return
	}

	app.ResponseSuccess(c, fmt.Sprintf("Token: %s", token))
}
