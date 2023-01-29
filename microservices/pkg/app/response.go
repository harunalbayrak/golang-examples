package app

import (
	"examples/microservices/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ResponseWithError(c *gin.Context, httpError int, errorCode int) {
	log.WithFields(log.Fields{
		"code":     httpError,
		"errorMsg": e.GetMsg(errorCode),
	}).Info("Unsuccessful response")

	c.AbortWithStatusJSON(httpError, gin.H{
		"error": e.GetMsg(errorCode),
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	log.WithFields(log.Fields{
		"code": http.StatusOK,
		"data": data,
	}).Info("Successful response")

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
