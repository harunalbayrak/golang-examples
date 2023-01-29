package controllers

import (
	"errors"
	"examples/microservices/models"
	"examples/microservices/pkg/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there is no token"})
		return
	}

	user_id, err := util.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func CurrentUserAccess(c *gin.Context) bool {
	user_id, err := util.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	user_id_str := strconv.FormatUint(uint64(user_id), 10)
	id := c.Params.ByName("id")

	if id == user_id_str {
		return true
	}

	user, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return errors.New("Unauthorized")
	}

	return nil
}
