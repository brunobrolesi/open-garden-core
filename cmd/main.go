package main

import (
	"log"

	"github.com/brunobrolesi/open-garden-core/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.StartApp()
}
