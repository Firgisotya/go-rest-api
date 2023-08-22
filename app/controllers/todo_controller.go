package controllers

import (
	"net/http"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
)


func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.ShouldBindJSON(&todo)

	db := config.DB
	result := db.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var todos []models.Todo
	db := config.DB
	db.Where("user_id = ?", userID).Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	todoID := c.Param("id")

	var todo models.Todo
	db := config.DB
	result := db.Where("id = ?", todoID).First(&todo)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	var updatedTodo models.Todo
	c.ShouldBindJSON(&updatedTodo)

	db.Model(&todo).Updates(updatedTodo)

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	todoID := c.Param("id")

	var todo models.Todo
	db := config.DB
	result := db.Where("id = ?", todoID).Delete(&todo)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
