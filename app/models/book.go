package models

import (
	"time"
)

type Books struct {
	ID         uint  `gorm:"primary_key:auto_increment" json:"id"` //
	Category   Category
	CategoryID uint `gorm:"foreignkey:CategoryID"`
	Title      string
	Author     string
	Year       string
	Stock      int
	Price      int
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
