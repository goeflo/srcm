package main

import (
	"log"

	"github.com/floriwan/srcm/backend"
	"github.com/floriwan/srcm/pkg/config"
)

func main() {

	// load config
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load configation file", err)
	}

	// start backend
	backend := backend.Initialize()
	backend.Run()

}
