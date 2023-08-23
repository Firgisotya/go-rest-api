package controllers

import (
	"net/http"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type inputCategory struct {
	Name string `json:"name" binding:"required"`
}

var (
	validateCategory = validator.New()
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

func CreateCategory(c *gin.Context) {
	var inputCategory inputCategory
	if err := c.ShouldBindJSON(&inputCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateCategory.Struct(inputCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	category.ID = uuid.New().String()
	category.Name = inputCategory.Name

	db := config.DB
	result := db.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	db := config.DB
	result := db.First(&category, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching category"})
		return
	}

	var inputCategory inputCategory
	if err := c.ShouldBindJSON(&inputCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateCategory.Struct(inputCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = inputCategory.Name

	result = db.Save(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	db := config.DB
	result := db.First(&category, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching category"})
		return
	}

	result = db.Delete(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting category"})
		return
	}

	c.JSON(http.StatusOK, category)
}
