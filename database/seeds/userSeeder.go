package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"golang.org/x/crypto/bcrypt"
)

type seedUsers struct{}

func (s seedUsers) Run() error {

	db := config.DB

	// Seed users
	users := []models.User{
		{Username: "admin", Email: "admin@gmail.com", Password: "password", Role: 1},
		{Username: "user", Email: "user@gmail.com", Password: "password", Role: 2},
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
