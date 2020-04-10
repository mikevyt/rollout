package models

import (
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
