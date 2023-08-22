package models

import (
	"time"
)

type Books struct {
	ID        string    `gorm:"size:36;not null;unique;primary_key;" json:"id"`
	Category  Category
	CategoryID string   `gorm:"size:36;index;fieldNullable", json:"category_id"`
	Title     string
	Author     string
	Year      string
	CreatedAt time.Time
	UpdatedAt time.Time
}