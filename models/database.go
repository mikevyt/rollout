package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// InitDB initializes the database
func InitDB(dburi string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
		return err
	}

	db = client.Database("rollout")

	return nil
}

// CreateUser creates new User
func CreateUser(user *User) (err error) {
	_, err = db.Collection("users").InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}
