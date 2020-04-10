package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	m "github.com/mikevyt/rollout/models"
)

func main() {
	err := m.InitDB("mongodb://localhost:27017")
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
