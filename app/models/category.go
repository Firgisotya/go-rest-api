package models

import (
	"time"
)

type Category struct {
	ID        uint           `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}