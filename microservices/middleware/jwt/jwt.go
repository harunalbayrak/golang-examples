package middleware

import (
	"examples/microservices/pkg/app"
	"examples/microservices/pkg/e"
	"examples/microservices/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c)
		if err != nil {
			app.ResponseWithError(c, http.StatusUnauthorized, e.ERROR_STATUS_UNAUTHORIZED)
			return
		}
		c.Next()
	}
}
