package main

import (
	"github.com/tomasff/brevis/internal/models"
	"github.com/tomasff/brevis/internal/routes"
	"log"
	"os"
)

func main() {
	log.SetPrefix("brevis - ")

	dbName := os.Getenv("DB_NAME")
	dbUri := os.Getenv("DB_URI")

	db, err := models.NewDB(dbName, dbUri)

	if err != nil {
		log.Fatalln(err)
	}

	r := routes.SetupRouter(db)

	err = r.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
