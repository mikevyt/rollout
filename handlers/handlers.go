package handlers

import (
	"net/http"

	m "github.com/mikevyt/rollout/models"
)

// UsersIndex GETs all Users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	user := m.User{Username: "donkykong"}
	m.CreateUser(&user)
	// users := m.Users{
	// 	m.User{ID: 1},
	// 	m.User{ID: 2},
	// }

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(users); err != nil {
	// 	panic(err)
	// }
}
