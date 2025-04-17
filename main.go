package main

import (
	"auth0_example/platform/authenticator"
	"auth0_example/platform/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	router := routes.New(auth)
	log.Println("Server listening on http://localhost:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", router); err != nil {
		panic(err)
	}
}
