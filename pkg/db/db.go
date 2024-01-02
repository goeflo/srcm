package db

import (
	"log"

	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db/model"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Initialize() {
	db, err := gorm.Open(sqlite.Open(config.GlobalConfig.DbSqliteFilename), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Instance = db
}

func Migrate() {
	if err := Instance.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}
	log.Println("db migration complete")
}
