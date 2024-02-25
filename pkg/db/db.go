package db

import (
	"log"

	"github.com/floriwan/srcm/pkg/db/model"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type SqlStorage struct {
	DB *gorm.DB
}

type SqlLiteConfig struct {
	Filename string
}

func NewSqlLiteStorage(cfg SqlLiteConfig) *SqlStorage {
	db, err := gorm.Open(sqlite.Open(cfg.Filename), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &SqlStorage{DB: db}
}

func (db *SqlStorage) PolulateInitialData() {
	res := db.DB.Where("email = ?", "admin").First(&model.User{})
	if res.Error != nil {
		log.Printf("create initial db data ...")
		// create admin user
		user := &model.User{Email: "admin", Admin: true}
		user.HashPassword("1234")

		res := db.DB.Create(&user)
		if res.Error != nil {
			log.Fatalf("error creating initial admin user: %v\n", res.Error)
		}

		// create user
		user = &model.User{Email: "user", Admin: false, SteamID: "0815"}
		user.HashPassword("1234")
		res = db.DB.Create(&user)
		if res.Error != nil {
			log.Fatalf("error creating initial admin user: %v\n", res.Error)
		}

		log.Printf("rows effected: %v\n", res.RowsAffected)
	}

}

/*
var Instance *gorm.DB

func Initialize() {
	db, err := gorm.Open(sqlite.Open(config.GlobalConfig.DbSqliteFilename), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Instance = db
}

func PolulateInitialData() {

	res := Instance.Where("email = ?", "admin").First(&model.User{})
	if res.Error != nil {
		log.Printf("create initial db data ...")
		// create admin user
		user := &model.User{Email: "admin", Admin: true}
		user.HashPassword("1234")

		res := Instance.Create(&user)
		if res.Error != nil {
			log.Fatalf("error creating initial admin user: %v\n", res.Error)
		}

		// create user
		user = &model.User{Email: "user", Admin: false, SteamID: "0815"}
		user.HashPassword("1234")
		res = Instance.Create(&user)
		if res.Error != nil {
			log.Fatalf("error creating initial admin user: %v\n", res.Error)
		}

		log.Printf("rows effected: %v\n", res.RowsAffected)
	}

}
*/
