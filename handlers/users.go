package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mikevyt/rollout/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUsers GETs all Users
// TODO: Handle Requests and Responses better
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}
	filter := bson.D{primitive.E{}}
	users, err := db.ReadUser(filter)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

// GetUser GETs a single User
func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	discordid, err := strconv.ParseInt(vars["discordid"], 10, 64)
	if err != nil {
		panic(err)
	}

	filter := bson.D{primitive.E{Key: "discordid", Value: discordid}}
	user, err := db.ReadUser(filter)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

// PostUserRequest defines a new User request
type PostUserRequest struct {
	DiscordID       int64
	DiscordUsername string
}

// PostUser POSTs a new User
func PostUser(w http.ResponseWriter, r *http.Request) {
	var request PostUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	user := models.NewUser(request.DiscordID, request.DiscordUsername)

	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}

	err = db.CreateUser(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// PutUserRequest defines an update User request
// TODO: Determine what changes can be made
type PutUserRequest struct {
}

// PutUser PUTs changes to a User
// TODO: Find out how to handle multiple changes
func PutUser(w http.ResponseWriter, r *http.Request) {
	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	discordid, err := strconv.ParseInt(vars["discordid"], 10, 64)
	if err != nil {
		panic(err)
	}

	filter := bson.D{primitive.E{Key: "discordid", Value: discordid}}
	update := bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{
			primitive.E{Key: "isdeleted", Value: true},
			primitive.E{Key: "updatedate", Value: time.Now()},
		},
	}}
	err = db.UpdateUser(filter, update)

	if err != nil {
		panic(err)
	}
}

// DeleteUser DELETEs a User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	discordid, err := strconv.ParseInt(vars["discordid"], 10, 64)
	if err != nil {
		panic(err)
	}

	filter := bson.D{primitive.E{Key: "discordid", Value: discordid}}
	update := bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{
			primitive.E{Key: "isdeleted", Value: true},
			primitive.E{Key: "updatedate", Value: time.Now()},
		},
	}}
	err = db.UpdateUser(filter, update)

	if err != nil {
		panic(err)
	}
}
