package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"gorm.io/gorm"
)

// seedBook
func seedBooks(db *gorm.DB) error {


	books := []models.Books{
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", CategoryID: 4, Year: "2022"},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Philosopher's Stone", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Chamber of Secrets", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Prisoner of Azkaban", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Goblet of Fire", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Order of the Phoenix", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Half-Blood Prince", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
		{Title: "Harry Potter and the Deathly Hallows", Author: "J.K. Rowling", CategoryID: 4, Year: "2022"},
	}

	for _, book := range books {
		if err := db.Create(&book).Error; err != nil {
			panic(err)
		}
	}

	return nil

}