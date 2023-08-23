package controllers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type inputBook struct {
	CategoryID string `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	Year       string `json:"year" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Image      string `json:"image" binding:"required"`
}

var (
	validateBooks = validator.New()
)

func GetAllBook(c *gin.Context) {
	var book []models.Books
	db := config.DB
	result := db.Find(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching books"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {

	// Print request body
    requestBody, _ := c.GetRawData()
    log.Printf("Request Body: %s\n", requestBody)

	var inputBook inputBook


	if err := c.ShouldBindJSON(&inputBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": err})
		return
	}

	if err := validateBooks.Struct(inputBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "body": err})
		return
	}

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	// Save the image to the uploads directory
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}


	var book models.Books
	book.ID = uuid.New().String()
	book.CategoryID = inputBook.CategoryID
	book.Title = inputBook.Title
	book.Author = inputBook.Author
	book.Year = inputBook.Year
	book.Stock = inputBook.Stock
	book.Price = inputBook.Price
	book.Image = filePath

	db := config.DB
	result := db.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func ShowBook(c *gin.Context) {
	var book models.Books
	db := config.DB
	result := db.First(&book, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Books
	db := config.DB
	result := db.First(&book, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching book"})
		return
	}

	var inputBook inputBook
	if err := c.ShouldBindJSON(&inputBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateBooks.Struct(inputBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.CategoryID = inputBook.CategoryID
	book.Title = inputBook.Title
	book.Author = inputBook.Author
	book.Year = inputBook.Year
	book.Stock = inputBook.Stock
	book.Price = inputBook.Price
	book.Image = inputBook.Image

	result = db.Save(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {

}
