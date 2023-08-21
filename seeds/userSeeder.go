package main

import (
	"fmt"

	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	db, err := gorm.Open(mysql.Open(config.GetDatabaseURL()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	config.SetDB(db)

	seedUsers()
}

func seedUsers() {
	db := config.GetDB()

	// Check if there are existing users
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		fmt.Println("Users already seeded.")
		return
	}

	// Seed users
	users := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
	}

	for _, user := range users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		db.Create(&user)
	}

	fmt.Println("Users seeded successfully.")
}
