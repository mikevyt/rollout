package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// User Model
type User struct {
	DiscordID       int64     `json:"discordId"`
	DiscordUsername string    `json:"discordUserName"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	IsDeleted       bool      `json:"isDeleted"`
}

// Users Model
type Users []User

func NewUser(discordid int64, discordusername string) *User {
	user := User{
		DiscordID:       discordid,
		DiscordUsername: discordusername,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return &user
}

func (u *User) SetUpdatedAt() {
	u.CreatedAt = time.Now()
}

func (db *DB) CreateUser(user *User) (err error) {
	_, err = db.Collection("users").InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ReadUser(filter bson.D) (*Users, error) {
	cur, err := db.Collection("users").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users Users
	for cur.Next(context.Background()) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return &users, nil
}

func (db *DB) UpdateUser(filter bson.D, update bson.D) error {
	_, err := db.Collection("users").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}
