package command

import (
	"fmt"
	"log"

	"github.com/Firgisotya/go-rest-api/app"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/Firgisotya/go-rest-api/database/seeds"
)

func Migrate() {
	fmt.Println("Running migration...")

	config.ConnectDB()

	for _, model := range app.RegisterModel(){
		// Drop Tabel yang sudah ada
		err := config.DB.Migrator().DropTable(model.Model)
		if err != nil {
			fmt.Printf("Failed to drop table for model %T: %v\n", model.Model, err)
		}


		// migrasi model baru
		err = config.DB.AutoMigrate(model.Model)

		if err != nil {
			panic("Failed to migrate database")
		}
	}

	fmt.Println("Migration completed")
}

func Seed(){

	config.ConnectDB()

	fmt.Println("Running seed...")
	
	if err := seeds.DBSeed(); err != nil {
		log.Fatal("Failed to seed: %v", err)
	}

	fmt.Println("Seed completed!")
}
