package models

import (
	"time"
)

type Category struct {
	ID        string          `gorm:"size:36;not null;unique;primary_key;" json:"id"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}