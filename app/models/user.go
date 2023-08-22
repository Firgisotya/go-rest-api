package models

import (
	"time"
)

type User struct {
	ID        string      `gorm:"size:36;not null;unique;primary_key;" json:"id"`
	Username string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

