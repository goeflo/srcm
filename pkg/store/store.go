package store

import (
	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/db/model"
)

type Storage struct {
	db *db.SqlStorage
}

type Store interface {
	CreateUser(u *model.User)
	GetUser(name string)
	CreateSeasons(name string)
	GetSeasons()
}

func NewStore(db *db.SqlStorage) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser(u model.User) error {
	return nil
}

func (s *Storage) GetUser(email string) (*model.User, error) {
	user := model.User{}
	res := s.db.DB.Where(&model.User{Email: email, Active: true}).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (s *Storage) GetSeasons() ([]model.Season, error) {
	seasons := []model.Season{}
	res := s.db.DB.Find(&seasons)
	if res.Error != nil {
		return seasons, res.Error
	}
	return seasons, nil
}

func (s *Storage) CreateSeason(name string) (*model.Season, error) {
	season := model.Season{Name: name}
	res := s.db.DB.Create(&season)
	if res.Error != nil {
		return nil, res.Error
	}
	return &season, nil
}
