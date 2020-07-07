package main

import (
	"github.com/tomasff/brevis/internal/routes"
	"log"
)

func main() {
	r := routes.SetupRouter()

	err := r.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
