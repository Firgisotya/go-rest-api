package main

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	db, err := gorm.Open(mysql.Open(config.GetDatabaseURL()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	db.AutoMigrate(&models.User{}, &models.Todo{})
	
}
