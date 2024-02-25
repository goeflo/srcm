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
	SeasonID int
	Name     string `gorm:"type:varchar(50);unique;not null" json:"name"`
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
