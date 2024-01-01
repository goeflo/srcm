package db

import (
	"log"

	"github.com/floriwan/srcm/pkg/db/model"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Initialize() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Instance = db
}

func Migrate() {
	Instance.AutoMigrate(&model.User{})
	log.Println("Database Migration Completed!")
}
