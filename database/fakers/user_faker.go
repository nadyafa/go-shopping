package fakers

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/nadyafa/go-shopping/app/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	gofakeit.Seed(0)

	return &models.User{
		ID:            gofakeit.UUID(),
		Addresses:     nil,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Email:         gofakeit.Email(),
		Password:      gofakeit.Password(true, true, true, false, false, 8),
		RememberToken: gofakeit.UUID(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
	}
}
