package model

import (
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

type Race struct {
	gorm.Model
	SeasonID uint
	Name     string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

type Driver struct {
	gorm.Model
	UserID    uint
	RaceID    uint
	CarNumber uint   `gorm:";unique;not null" json:"carNumber"`
	Name      string `gorm:"type:varchar(50);unique;not null" json:"name"`
	TeamName  string `gorm:"type:varchar(50);unique;not null" json:"teamName"`
}

type RaceResult struct {
	gorm.Model
	RaceID           int
	Pos              uint
	StartPos         uint
	Driver           string `gorm:"type:varchar(50)" json:"driver"`
	Team             string
	Car              string
	Class            string
	TotalTime        uint
	BestLapTime      uint
	BestCleanLapTime uint
	CalculatedPoints uint
	FinalPoints      uint
	Laps             uint
}
