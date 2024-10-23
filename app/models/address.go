package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User       User
	UserID     string `gorm:"size:36;index"`
	Name       string `gorm:"size:100"`
	isPrimary  bool
	CityID     string `gorm:"size:100"`
	ProvinceID string `gorm:"size:100"`
	Address1   string `gorm:"size:255"`
	Address2   string `gorm:"size:255"`
	Phone      string `gorm:"size:100"`
	PostCode   string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}