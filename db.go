package main

import (
	"os"

	"github.com/JedBeom/onionpi/models"
	"github.com/go-pg/pg"
)

var (
	db *pg.DB
)

func ConnectDB() {
	db = pg.Connect(&pg.Options{
		User:     "onionpi",
		Password: "yummypi",
		Database: "onionpi",
	})

	if len(os.Args) > 1 {
		err := models.CreateTables(db)
		if err != nil {
			panic(err)
		}
	}
}
