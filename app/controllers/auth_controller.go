package controllers

import (
	"net/http"
	"time"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

var currentUser models.User // Simpan pengguna yang sedang login



func Register(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	user.Password = string(hashedPassword)

	db := config.GetDB()
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
	db := config.GetDB()
	result := db.Where("username = ?", inputUser.Username).First(&dbUser)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(inputUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"user_id":  dbUser.ID,
	})

	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	currentUser = dbUser

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "username": dbUser.Username})
}

