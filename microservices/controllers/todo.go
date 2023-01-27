package controllers

import (
	"examples/microservices/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	userID := c.Params.ByName("id")

	var todo []models.Todo
	err := models.GetAllTodos(&todo, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")

	var todo models.Todo
	err := models.GetATodo(&todo, todoID, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")

	var todo models.Todo
	err := models.GetATodo(&todo, todoID, userID)
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
	userID := c.Params.ByName("id")
	todoID := c.Params.ByName("todo_id")

	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.DeleteATodo(&todo, todoID, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
