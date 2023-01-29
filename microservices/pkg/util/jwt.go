package util

import (
	"errors"
	"examples/microservices/pkg/e"
	"examples/microservices/pkg/setting"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GenerateToken(user_id uint) (string, error) {
	expirationTime, err := GenerateExpirationTime()
	if err != nil {
		log.Error(e.GetMsg(e.ERROR_TOKEN_GENERATION_FAIL))
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = expirationTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Infof("New token is generated: %s", token)

	return token.SignedString([]byte(setting.AppSettings.TokenSettings.ApiSecretKey))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := ParseToken(tokenString)
	if err != nil {
		return err
	}
	log.Infof("Token is valid: %s", tokenString)

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error(e.GetMsg(e.ERROR_UNEXPECTED_SIGNING_METHOD))
			return nil, errors.New(e.GetMsg(e.ERROR_UNEXPECTED_SIGNING_METHOD))
		}
		return []byte(setting.AppSettings.TokenSettings.ApiSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateExpirationTime() (int64, error) {
	token_lifespan, err := strconv.Atoi(setting.AppSettings.TokenSettings.TokenHourLifespan)
	if err != nil {
		log.Error(e.GetMsg(e.ERROR_TOKENHOURLIFESPAN_CONVERTION))
		return 0, err
	}

	return time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(), nil
}
