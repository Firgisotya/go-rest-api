package controllers

import (
	"net/http"

	"github.com/Firgisotya/go-rest-api/app/middlewares"
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	user.Password = string(hashedPassword)

	db := config.DB	

	var existUser models.User

	// cek duplikasi email
	if err := config.DB.Where("email = ?", user.Email).First(&existUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exist"})
		return
	}
		// cek duplikasi username
	if err := config.DB.Where("username = ?", user.Username).First(&existUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username already exist"})
		return
	}

	
	
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var inputUser models.User
	c.ShouldBindJSON(&inputUser)

	var dbUser models.User
	db := config.DB

	//cek apakah username ada di database 
	result := db.Where("username = ?", inputUser.Username).First(&dbUser)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username"})
		return
	}

	// cek password
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(inputUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}

	// fungsi berhasil login
	token, err := middlewares.GenerateToken(dbUser.ID, dbUser.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error creating token"})
		return
	}

	// tampilkan data hasil login
	c.JSON(http.StatusOK, gin.H{
		"Token-type": "Bearer",
		"Token": token,
		"User": dbUser,
	})

	
}

func Logout(c *gin.Context){
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"message": "Logout success"})
}
