package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Lastname         string  `gorm:"type:varchar(50);not null" json:"lastname"`
	Password         string  `gorm:"type:varchar(50);not null;default:null" json:"password"`
	Email            string  `gorm:"type:varchar(50);unique;not null;default:null" json:"email"`
	SteamID          string  `gorm:"type:varchar(20);not null" json:"steamID"`
	VerificationCode *string `json:"verificationCode"`
	Verified         bool    `gorm:"not null" json:"verified"`
	Active           bool    `gorm:"type:bool;default:true;not null" json:"active"`
	Admin            bool    `gorm:"type:bool;default:false;not null" json:"admin"`
}

type UserToEvent struct {
	UserID  int64
	EventID int64
	Car     Car
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
