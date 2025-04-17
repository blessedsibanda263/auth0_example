package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println(os.Getenv("AUTH0_CALLBACK_URL"))
	log.Println(os.Getenv("AUTH0_DOMAIN"))
}
