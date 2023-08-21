package models

import (
	"time"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

