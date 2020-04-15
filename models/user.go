package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// DiscordUser Model
type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int64  `json:"public_flags"`
	Flags         int64  `json:"flags"`
	Locale        string `json:"locale"`
	MFAEnabled    bool   `json:"mfa_enabled"`
}

// User Model
type User struct {
	DiscordUser
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
}

// Users Model
type Users []User

// NewUser creates new user
func NewUser(discordUserData DiscordUser) *User {
	user := User{
		DiscordUser: discordUserData,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &user
}

// CreateUser adds User to database
func (db *DB) CreateUser(user *User) (err error) {
	_, err = db.Collection("users").InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}

// ReadUser reads User from database
func (db *DB) ReadUser(filter bson.D) (*User, error) {
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

	return &users[0], nil // TODO: fix this
}

// UpdateUser updates User in database
func (db *DB) UpdateUser(filter bson.D, update bson.D) error {
	_, err := db.Collection("users").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}
