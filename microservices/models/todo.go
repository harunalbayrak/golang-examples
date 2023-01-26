package models

import(
	"fmt"
	"examples/microservices/config"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID uint            `json:"id"`
	Title string       `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAllTodos(todos *[]Todo) (err error) {
	if err = config.DB.Find(todos).Error; err != nil {
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

func GetATodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

func DeleteATodo(todo *Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
