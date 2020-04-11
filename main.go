package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mikevyt/rollout/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	err = models.StartDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := NewRouter()

	fmt.Printf("Server Running. Listening at: localhost:%s.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
