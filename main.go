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

	dbConfig := db.SqlLiteConfig{
		Filename: config.GlobalConfig.DbSqliteFilename,
	}

	sqlStorage := db.NewSqlLiteStorage(dbConfig)
	sqlStorage.PolulateInitialData()
	store := store.NewStore(sqlStorage)
	server := api.NewAPIServer(config.GlobalConfig.HttpPort, *store)
	server.Serve()

	// start backend
	//backend := backend.Initialize()
	//backend.Run()

}
