package seeds

import (

	"github.com/Firgisotya/go-rest-api/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)



func seedUsers(db *gorm.DB) error {
	// Seed users
	users := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
	}

	for _, user := range users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
