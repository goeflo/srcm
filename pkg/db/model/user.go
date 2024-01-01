package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name             string    `gorm:"type:varchar(255);not null"`
	Lastname         string    `json:"lastname" binding:"required"`
	Password         string    `gorm:"not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	SteamID          string    `gorm:"not null"`
	VerificationCode string
	Verified         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword)); err != nil {
		return err
	}
	return nil
}
