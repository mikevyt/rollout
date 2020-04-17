package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mikevyt/rollout/models"
	"github.com/mitchellh/mapstructure"
)

// DiscordCallback handles retrieving user information from Discord
func DiscordCallback(w http.ResponseWriter, r *http.Request) {
	userResponse, err := gothic.CompleteUserAuth(w, r)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	discordUser := models.DiscordUser{}
	mapstructure.Decode(userResponse.RawData, &discordUser)

	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}

	user := models.NewUser(discordUser)
	err = db.CreateUser(user)

	if err := json.NewEncoder(w).Encode(discordUser); err != nil {
		panic(err)
	}
}

// Authenticate is the entry point for authenticating with Discord
func Authenticate(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	fmt.Println(gothUser)
	if err == nil {
		fmt.Println("User found")
		if err := json.NewEncoder(w).Encode(gothUser); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Beginning Auth Handler")
		gothic.BeginAuthHandler(w, r)
	}
}

// Logout logs out the user
func Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
