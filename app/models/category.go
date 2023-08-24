package models

import (
	"time"
)

type Category struct {
	ID       uint `gorm:"primary_key:auto_increment" json:"id"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}