package store

import (
	"fmt"
	"log"

	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/db/model"
)

type Store interface {
	AddUser(u *model.User) error
	GetUser(email string) (*model.User, error)

	AddSeason(name string) (*model.Season, error)
	GetSeasons() ([]model.Season, error)
	GetSeasonByID(ID uint) (*model.Season, error)

	AddRace(raceName string, seasonID uint) (*model.Race, error)
	GetRaceByID(ID uint) (*model.Race, error)

	AddDriver(driverName string, teamName string, raceID uint) (*model.Driver, error)
}

type Storage struct {
	db *db.SqlLiteDB
}

func NewStorage(db *db.SqlLiteDB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) AddUser(u *model.User) error {
	return nil
}

func (s *Storage) GetUser(email string) (*model.User, error) {
	user := model.User{}
	if err := s.db.DB.Where(&model.User{Email: email, Active: true}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) GetSeasons() ([]model.Season, error) {
	seasons := []model.Season{}
	if err := s.db.DB.Find(&seasons).Error; err != nil {
		return seasons, err
	}
	return seasons, nil
}

func (s *Storage) AddSeason(name string) (*model.Season, error) {
	log.Printf("CreateSeason name:%v\n", name)
	season := model.Season{Name: name}
	if err := s.db.DB.Create(&season).Error; err != nil {
		return nil, err
	}
	return &season, nil
}

func (s *Storage) GetSeasonByID(ID uint) (*model.Season, error) {
	log.Printf("GetSeasonByID ID:%v\n", ID)
	season := model.Season{}
	if err := s.db.DB.First(&season, ID).Error; err != nil {
		return nil, err
	}
	return &season, nil
}

func (s *Storage) GetSeason(name string) (*model.Season, error) {
	log.Printf("GetSeason name:%v\n", name)
	season := model.Season{}
	if err := s.db.DB.Where(&model.Season{Name: name}).First(&season).Error; err != nil {
		return nil, err
	}
	return &season, nil
}

func (s *Storage) AddRace(raceName string, seasonID uint) (*model.Race, error) {
	if raceName == "" || seasonID == 0 {
		return nil, fmt.Errorf("race name and season name must be set")
	}

	race := model.Race{Name: raceName}
	log.Printf("create race '%v' in season ID '%v'\n", raceName, seasonID)

	season, err := s.GetSeasonByID(seasonID)
	if err != nil {
		return nil, err
	}

	race.SeasonID = season.ID

	res := s.db.DB.Create(&race)
	if res.Error != nil {
		return nil, res.Error
	}
	return &race, nil

}

func (s *Storage) GetRaceByID(ID uint) (*model.Race, error) {
	log.Printf("GetRaceByID ID:%v\n", ID)
	race := model.Race{}
	if err := s.db.DB.First(&race, ID).Error; err != nil {
		return nil, err
	}
	return &race, nil
}

func (s *Storage) AddDriver(driverName string, teamName string, raceID uint) (*model.Driver, error) {
	if driverName == "" {
		return nil, fmt.Errorf("driver name must be set")
	}

	race, err := s.GetRaceByID(raceID)
	if err != nil {
		return nil, err
	}

	driver := model.Driver{Name: driverName, TeamName: teamName, RaceID: race.ID}
	res := s.db.DB.Create(&driver)
	if res.Error != nil {
		return nil, res.Error
	}
	return &driver, nil
}
