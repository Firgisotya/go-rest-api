package main

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/app/routes"
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

	config.SetDB(db)

	db.AutoMigrate(&models.User{}, &models.Todo{})
	
	r := routes.SetupRouter()
	r.Run() // Menjalankan server pada port 8080
}
