package models

import (
	"errors"
	"examples/microservices/config"
	"math/rand"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Type     string `json:"type"`
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
	if err := config.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAUser(user *User, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}

	return nil
}

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

func VerifyPassword(userPassword string, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if err := config.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil
}

func CountUsers() int64 {
	var count int64

	config.DB.Table("users").Count(&count)

	return count
}

func GetRandomUserId() int {
	var u User
	countUser := CountUsers()
	offset := rand.Intn(int(countUser))
	config.DB.Offset(offset).Find(&u)

	return int(u.ID)
}
