package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/google/uuid"
)

type seedCategories struct{}

func (s seedCategories) Run() error {

	db := config.DB

	//seed categories
	categories := []models.Category{
		{ID: uuid.New().String(), Name: "Novel"},
		{ID: uuid.New().String(), Name: "Comic"},
		{ID: uuid.New().String(), Name: "Science"},
		{ID: uuid.New().String(), Name: "History"},
		{ID: uuid.New().String(), Name: "Programming"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			panic(err)
		}

	}
	return nil

}

