package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mikevyt/rollout/auth"
	"github.com/mikevyt/rollout/filters"
	"github.com/mikevyt/rollout/models"
)

// DiscordOauth redirects to Discord's Auth Page
func DiscordOauth(w http.ResponseWriter, r *http.Request) {
	discordLogin := auth.GetDiscordAuthURL()
	http.Redirect(w, r, discordLogin, http.StatusSeeOther)
}

// Login handles retrieving user information from Discord
// TODO: Improve Error handling here
func Login(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		panic(errors.New("No 'code' in query parameters"))
	}

	accessToken := auth.GetAccessToken(code)
	discordUserData := auth.GetUserData(accessToken)

	db, err := models.GetDB()
	if err != nil {
		panic(err)
	}

	filter := filters.GetUserByDiscordID(discordUserData.ID)

	user, err := db.ReadUser(filter)

	if user == nil {
		fmt.Println("new user")
		user = models.NewUser(discordUserData)

		err = db.CreateUser(user)

		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&user); err != nil {
			panic(err)
		}
	} else if user.DiscordUser != discordUserData {
		fmt.Println("update user")
		user.DiscordUser = discordUserData
		update := filters.UpdateUser(discordUserData)
		err = db.UpdateUser(filter, update)
	}

	if err := json.NewEncoder(w).Encode(accessToken); err != nil {
		panic(err)
	}
}
