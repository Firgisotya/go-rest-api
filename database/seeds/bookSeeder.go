package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/google/uuid"
)

type seedBooks struct{}

// seedBook
func (s seedBooks) Run() error {

	db := config.DB


	books := []models.Books{
		{ID: uuid.New().String(), Title: "The Lord of The Rings", Author: "J.R.R Tolkien", CategoryID: "", Year: "1945"},
		{ID: uuid.New().String(), Title: "The Hobbit", Author: "J.R.R Tolkien", CategoryID: "", Year: "1945"},
	}

	for _, book := range books {
		if err := db.Create(&book).Error; err != nil {
			panic(err)
		}
	}

	return nil

}