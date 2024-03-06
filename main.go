package main

import (
	"embed"
	"io/fs"
	"log"

	"github.com/floriwan/srcm/pkg/api"
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/store"

	_ "embed"
)

//go:embed public
var public embed.FS

func main() {

	// embed static html pages
	publicHtmlDir, _ := fs.Sub(public, "public")

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
	server := api.NewAPIServer(config.GlobalConfig.HttpPort, config.GlobalConfig.ApiPort, store, publicHtmlDir)
	server.Serve()

}
