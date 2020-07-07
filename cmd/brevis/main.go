package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	err := r.Run()

	if err != nil {
		log.Fatal(err)
	}
}
