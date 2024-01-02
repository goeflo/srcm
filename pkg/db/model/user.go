package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               int64  `gorm:"uniqueIndex,primaryKey,not null;default:null"`
	Name             string `gorm:"type:varchar(50);not null"`
	Lastname         string `gorm:"type:varchar(50);not null"`
	Password         string `gorm:"type:varchar(50);not null;default:null"`
	Email            string `gorm:"type:varchar(50);unique;not null;default:null"`
	SteamID          string `gorm:"type:varchar(20);not null"`
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
