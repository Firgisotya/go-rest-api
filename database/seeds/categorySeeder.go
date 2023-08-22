package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"gorm.io/gorm"
)


func seedCategories(db *gorm.DB) error {

	//seed categories
	categories := []models.Category{
		{Name: "Fiction"},
		{Name: "Non-Fiction"},
		{Name: "Science Fiction"},
		{Name: "Fantasy"},
		{Name: "Romance"},
		{Name: "Mystery"},
		{Name: "Horror"},
		{Name: "Thriller"},
		{Name: "Children"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			panic(err)
		}

	}
	return nil

}

