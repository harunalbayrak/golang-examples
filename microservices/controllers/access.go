package controllers

import (
	"errors"
	"examples/microservices/models"
	"examples/microservices/pkg/app"
	"examples/microservices/pkg/e"
	"examples/microservices/pkg/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_NOT_FOUND_TOKEN)
		return
	}

	user_id, err := util.ExtractTokenID(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_EXTRACT_TOKEN)
		return
	}

	user, err := models.GetUserByID(user_id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_NOT_FOUND_USER)
		return
	}

	app.ResponseSuccess(c, user)
}

func CurrentUserAccess(c *gin.Context) bool {
	user_id, err := util.ExtractTokenID(c)
	if err != nil {
		return false
	}

	user_id_str := strconv.FormatUint(uint64(user_id), 10)
	id := c.Params.ByName("id")

	if id == user_id_str {
		return true
	}

	user, err := models.GetUserByID(user_id)
	if err != nil {
		return false
	}

	if strings.EqualFold(user.Type, "Admin") {
		return true
	}

	return false
}

func CheckAccess(c *gin.Context) error {
	access := CurrentUserAccess(c)
	if !access {
		return errors.New("Unauthorized")
	}

	return nil
}
