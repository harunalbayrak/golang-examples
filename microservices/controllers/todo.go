package controllers

import (
	"examples/microservices/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	userID := c.Params.ByName("id")
	var todo []models.Todo
	err = models.GetAllTodos(&todo, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	userID := c.Params.ByName("id")
	userIDint, _ := strconv.Atoi(userID)
	var todo models.Todo
	c.BindJSON(&todo)
	todo.UserID = uint(userIDint)
	err = models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")
	var todo models.Todo
	err = models.GetATodo(&todo, todoID, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
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
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	err := CheckAccess(c)
	if err != nil {
		return
	}

	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")

	id := c.Params.ByName("id")
	err = models.DeleteATodo(todoID, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
