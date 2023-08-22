package controllers

import (
	"net/http"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
)


func GetAllCategory(c *gin.Context) {
		var category []models.Category
		db := config.DB
		result := db.Find(&category)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching categories"})
			return
		}

		c.JSON(http.StatusOK, category)
}