package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Lastname         string  `gorm:"type:varchar(50)" json:"lastname"`
	Firstname        string  `gorm:"type:varchar(50)" json:"firstname"`
	Password         string  `gorm:"type:varchar(50)" json:"password"`
	Email            string  `gorm:"type:varchar(50);uniqueIndex;not null;default:null" json:"email"`
	SteamID          string  `gorm:"type:varchar(20);not null" json:"steamID"`
	VerificationCode *string `json:"verificationCode"`
	Verified         bool    `gorm:"not null" json:"verified"`
	Active           bool    `gorm:"type:bool;default:true;not null" json:"active"`
	Admin            bool    `gorm:"type:bool;default:false;not null" json:"admin"`
	Participations   []Participation
}

type Participation struct {
	gorm.Model
	UserID  int `gorm:"not null"`
	EventID int `gorm:"not null"`
	Car     int `gorm:"not null"`
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

// func (user *User) GetByEmail(email string) error {
// 	//u := User{}
// 	res := db.Instance.Where("email = ?", email).First(&user)
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	if res.RowsAffected == 0 {
// 		return fmt.Errorf("no user with email %v found", email)
// 	}
// 	return nil
// }
