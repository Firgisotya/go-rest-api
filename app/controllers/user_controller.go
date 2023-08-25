package controllers

import (
	"net/http"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type inputUser struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role uint `json:"role" binding:"required"`
}

var (
	validateUser = validator.New()
)


func GetAllUser(c *gin.Context) {
	var user []models.User
	db := config.DB
	result := db.Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var inputUser inputUser
	db := config.DB
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateUser.Struct(inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	var user models.User
	
	if err := db.Where("username = ?", inputUser.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username already exist"})
		return
	}
	
	if err := db.Where("email = ?", inputUser.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exist"})
		return
	}

	user.Username = inputUser.Username
	user.Email = inputUser.Email
	user.Password = string(hashedPassword)
	user.Role = int(inputUser.Role)
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func ShowUser(c *gin.Context){
	var user models.User

	db := config.DB

	id := c.Param("id")
	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	db := config.DB
	id := c.Param("id")
	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": result.Error})
		return
	}

	var inputUser inputUser
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateUser.Struct(inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	user.Username = inputUser.Username
	user.Email = inputUser.Email
	user.Password = string(hashedPassword)
	user.Role = int(inputUser.Role)

	result = db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	db := config.DB
	id := c.Param("id")
	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user"})
		return
	}

	result = db.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}