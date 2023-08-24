package models

import (
	"time"
)

type User struct {
	ID       uint `gorm:"primary_key:auto_increment" json:"id"`
	Username string
	Email string
	Password string
	Role     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

