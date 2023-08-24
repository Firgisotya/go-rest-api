package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
)

type seedBooks struct{}

// seedBook
func (s seedBooks) Run() error {

	db := config.DB


	books := []models.Books{
		{Title: "The Lord of The Rings", Author: "J.R.R Tolkien", CategoryID: 1, Year: "1945"},
		{Title: "The Hobbit", Author: "J.R.R Tolkien", CategoryID: 2, Year: "1945"},
	}

	for _, book := range books {
		if err := db.Create(&book).Error; err != nil {
			panic(err)
		}
	}

	return nil

}