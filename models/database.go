package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	mongo.Database
}

var db DB

// InitDB initializes the database
func StartDB(dburi string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
		return err
	}

	rollout := client.Database("rollout")
	db = DB{*rollout}

	return nil
}

func GetDB() (*DB, error) {
	if db == (DB{}) {
		return nil, errors.New("Database not initialized.")
	}
	return &db, nil
}
