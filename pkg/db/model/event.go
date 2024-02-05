package model

import (
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	Name string
}

type Race struct {
	gorm.Model
	SeasonID int
	Name     string
}

type RaceResult struct {
	gorm.Model
	RaceID           int
	Pos              uint
	StartPos         uint
	Participant      string
	Car              string
	Class            string
	TotalTime        uint
	BestLapTime      uint
	BestCleanLapTime uint
	CalculatedPoints uint
	FinalPoints      uint
}
