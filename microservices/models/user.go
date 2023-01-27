package models

import (
	"errors"
	"examples/microservices/config"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model   `gorm:"primary_key" json:"id"`
	Username     string `gorm:"size:255;not null;unique" json:"username"`
	Password     string `gorm:"size:255;not null;" json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Type         string `json:"type"`
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func GetAllUsers(users *[]User) (err error) {
	if err = config.DB.Find(users).Error; err != nil {
		return err
	}

	return nil
}

func CreateAUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAUser(user *User, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAUser(user *User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

func DeleteAUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}

func VerifyPassword(userPassword string, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
}

// func CheckAuth(username string, password string) (err error) {

// }

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := config.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}
