package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mikevyt/rollout/auth"
	"github.com/mikevyt/rollout/models"
)

// DiscordRedirect redirects to Discord's Auth Page
func DiscordRedirect(w http.ResponseWriter, r *http.Request) {
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

	if true { // new user
		user := models.NewUser(discordUserData)

		err = db.CreateUser(user)

		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&user); err != nil {
			panic(err)
		}
	}
}
