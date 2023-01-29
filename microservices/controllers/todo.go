package controllers

import (
	"examples/microservices/models"
	"examples/microservices/pkg/app"
	"examples/microservices/pkg/e"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	userID := c.Params.ByName("id")
	var todo []models.Todo

	err = models.GetAllTodos(&todo, userID)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_GET_TODOS)
	} else {
		app.ResponseSuccess(c, todo)
	}
}

func CreateATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	userID := c.Params.ByName("id")
	userIDint, _ := strconv.Atoi(userID)
	var todo models.Todo
	c.BindJSON(&todo)
	todo.UserID = uint(userIDint)

	err = models.CreateATodo(&todo)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_CREATE_TODO)
	} else {
		app.ResponseSuccess(c, todo)
	}
}

func GetATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")
	var todo models.Todo

	err = models.GetATodo(&todo, todoID, userID)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_GET_TODO)
	} else {
		app.ResponseSuccess(c, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")
	var todo models.Todo
	err = models.GetATodo(&todo, todoID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)

	err = models.UpdateATodo(&todo)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_UPDATE_TODO)
	} else {
		app.ResponseSuccess(c, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_STATUS_UNAUTHORIZED)
		return
	}

	todoID := c.Params.ByName("todo_id")
	id := c.Params.ByName("id")

	err = models.DeleteATodo(todoID, id)
	if err != nil {
		app.ResponseWithError(c, http.StatusBadRequest, e.ERROR_DELETE_TODO)
	} else {
		app.ResponseSuccess(c, fmt.Sprintf("Deleted todo: %s", id))
	}
}
