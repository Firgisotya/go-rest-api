package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(database *gorm.DB) {
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetDatabaseURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_NAME"),
	)
}

func GetPort() string {
	return fmt.Sprintf(":%s", GetEnv("PORT"))
}

func GetSecretKey() string {
	return GetEnv("SECRET_KEY")
}

