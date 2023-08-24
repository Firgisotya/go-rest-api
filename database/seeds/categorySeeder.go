package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
)

type seedCategories struct{}

func (s seedCategories) Run() error {

	db := config.DB

	//seed categories
	categories := []models.Category{
		{Name: "Novel"},
		{Name: "Comic"},
		{Name: "Science"},
		{Name: "History"},
		{Name: "Programming"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			panic(err)
		}

	}
	return nil

}

