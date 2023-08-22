package models

import (
	"time"
)

type Books struct {
	ID        uint           `gorm:"primaryKey"`
	CategoryID uint
	Title     string
	Author     string
	Year      string
	CreatedAt time.Time
	UpdatedAt time.Time
}