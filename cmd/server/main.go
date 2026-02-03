package main

import (
	"log"

	"github.com/muresan-gabriel/go-satellite-pass-prediction-api/internal/api"
)

func main() {
	passHandler := api.NewPassHandler()
	router := api.NewRouter(passHandler)

	log.Println("Starting server on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}