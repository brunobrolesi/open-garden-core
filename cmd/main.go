package main

import (
	"log"
	"os"

	"github.com/brunobrolesi/open-garden-core/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	err = godotenv.Load(curDir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.StartApp()
}
