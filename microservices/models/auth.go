package models

import (
	"examples/microservices/config"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func CheckAuth(username string, password string) (uint, error) {
	var auth User
	err := config.DB.Model(User{}).Where("username = ?", username).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound && auth.Username == "" {
		return 0, err
	}

	if auth.Username == "" {
		return 0, err
	}

	if auth.ID > 0 {
		err := VerifyPassword(password, auth.Password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return 0, err
		}

		return auth.ID, nil
	}

	return 0, nil
}
