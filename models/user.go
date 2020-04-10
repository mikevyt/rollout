package models

import (
	"context"
	"time"
)

// User Model
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Users Model
type Users []User

func NewUser(username string) *User {
	user := User{
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
