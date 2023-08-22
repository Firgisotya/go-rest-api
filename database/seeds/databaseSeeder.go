package seeds

import (
	"fmt"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeed(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: seedUsers(db)},
		{Seeder: seedCategories(db)},
		{Seeder: seedBooks(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	fmt.Println("Running seed...")
	for _, seeder := range RegisterSeed(db){
		err := db.Debug().Create(seeder.Seeder).Error

		if err != nil {
			return err
		}
	}
	return nil
}