package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mikevyt/rollout/models"
)

type PostUserRequest struct {
	Username string
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var request PostUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	user := models.NewUser(request.Username)

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
