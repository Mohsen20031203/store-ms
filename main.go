package main

import (
	"hello/config"
	"hello/internal/api"
	"hello/internal/db"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		return
	}
	db, err := db.NewStorege(config)
	if err != nil {
		return
	}

	server, err := api.NewServer(db, &config)
	if err != nil {
		return
	}
	server.Router.Run(":9091")
}
