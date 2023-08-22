package seeds

import (
	"github.com/Firgisotya/go-rest-api/app/models"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type seedUsers struct{}

func (s seedUsers) Run() error {

	db := config.DB

	// Seed users
	users := []models.User{
		{ID: uuid.New().String(), Username: "user1", Password: "password1"},
		{ID: uuid.New().String(), Username: "user2", Password: "password2"},
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
