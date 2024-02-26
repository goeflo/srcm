package main

import (
	"log"

	"github.com/floriwan/srcm/pkg/api"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/store"
)

func main() {

	// load config
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	// create slqLite database
	dbConfig := db.SqlLiteConfig{
		Filename: config.GlobalConfig.DbSqliteFilename,
	}

	sqlDB := db.NewSqlLiteDB(dbConfig)
	sqlDB.PolulateInitialData()

	// start api server
	store := store.NewStorage(sqlDB)
	server := api.NewAPIServer(config.GlobalConfig.HttpPort, store)
	server.Serve()

}
