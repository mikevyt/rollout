package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mikevyt/rollout/auth"
)

// DiscordRedirect redirects to Discord's Auth Page
func DiscordRedirect(w http.ResponseWriter, r *http.Request) {
	discordlogin := auth.GetDiscordAuthURL()
	http.Redirect(w, r, discordlogin, http.StatusSeeOther)
}

// Login handles retrieving user information from Discord
// TODO: Improve Error handling here
func Login(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		panic(errors.New("No 'code' in query parameters"))
	}

	accesstoken := auth.GetAccessToken(code)
	discordUserData := auth.GetUserData(accesstoken)

	if err := json.NewEncoder(w).Encode(&discordUserData); err != nil {
		panic(err)
	}
}
