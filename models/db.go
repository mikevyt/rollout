package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB contains a MongoDB instance
type DB struct {
	mongo.Database
}

var db DB

// StartDB initializes the MongoDB database
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

// GetDB gets instance of Mongo database
func GetDB() (*DB, error) {
	if db == (DB{}) {
		return nil, errors.New("database not initialized")
	}
	return &db, nil
}
