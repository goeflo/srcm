package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name       string
	SeriesID   int64
	Track      Track
	CarGroup   CarGroup
	RaceLength time.Time
}

type Season struct {
	gorm.Model
	Events []Event
}

type Car int
type CarGroup int
type Track int

const (
	Porsche_GT3 = iota
	Ferrari_GT3
)

const (
	All = iota
	GT3
	GT4
	GTC
	TCX
)

const (
	Barcelona = iota
	Bathurst
	Imola
	Monza
)
