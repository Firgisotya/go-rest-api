package command

import (
	"fmt"

	"github.com/Firgisotya/go-rest-api/app"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/Firgisotya/go-rest-api/database/seeds"
	"gorm.io/gorm"
)

func Migrate() {
	fmt.Println("Running migration...")

	config.ConnectDB()

	for _, model := range app.RegisterModel(){
		err := config.DB.AutoMigrate(model.Model)

		if err != nil {
			panic("Failed to migrate database")
		}
	}

	fmt.Println("Migration completed")
}

func Seed(db *gorm.DB) error{
	fmt.Println("Running seed...")
	for _, seeder := range seeds.RegisterSeed(db){
		err := db.Debug().Create(seeder.Seeder).Error

		if err != nil {
			return err
		}
	}
	return nil
}
