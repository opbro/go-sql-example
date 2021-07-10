package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "xadmin",
		Password: "xadmin",
		Addr:     "localhost:5432",
		Database: "people",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Println("Failed to Connect")

		os.Exit(100)
	}
	log.Println("Connected to Databse.")
	return db

}
