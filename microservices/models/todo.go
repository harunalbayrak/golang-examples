package models

import (
	"examples/microservices/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAllTodos(todos *[]Todo, userID string) (err error) {
	if err = config.DB.Where("user_id = ?", userID).Find(todos).Error; err != nil {
		return err
	}

	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id string, userID string) (err error) {
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo) (err error) {
	if err := config.DB.Save(todo).Error; err != nil {
		return err
	}

	return nil
}

func DeleteATodo(id string, userID string) (err error) {
	var todo Todo
	GetATodo(&todo, id, userID)
	if err := config.DB.Delete(todo).Error; err != nil {
		return err
	}

	return nil
}
