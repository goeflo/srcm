package migration

import (
	"log"

	"github.com/floriwan/srcm/pkg/db/model"
	"gorm.io/gorm"
)

type Migrator struct {
	DB *gorm.DB
}

func NewMigrator(db *gorm.DB) Migrator {
	return Migrator{DB: db}
}

func (m *Migrator) Migration() {
	if err := m.DB.AutoMigrate(
		&model.User{},
		&model.Participation{},
		&model.Season{},
		&model.Race{},
		&model.RaceResult{},
	); err != nil {
		log.Fatal(err)
	}
	log.Println("db migration complete")
}
