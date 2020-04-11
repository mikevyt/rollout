package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const authorizeURL = "https://discordapp.com/api/oauth2/authorize"
const tokenURL = "https://discordapp.com/api/oauth2/token"
const userDataURL = "https://discordapp.com/api/users/@me"
const redirectURL = "http://127.0.0.1:8080/login"
const scope = "identify email"

// GetDiscordAuthURL gets discord login URL.
func GetDiscordAuthURL() string {
	authURL, err := url.Parse(authorizeURL)
	if err != nil {
		panic(err)
	}

	clientID := os.Getenv("DISCORD_CLIENT_ID")

	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("redirect_uri", redirectURL)
	params.Add("response_type", "code")
	params.Add("scope", scope)

	authURL.RawQuery = params.Encode()

	return authURL.String()
}

type getAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// GetAccessToken gets the access token from Discord
func GetAccessToken(code string) string {
	clientID := os.Getenv("DISCORD_CLIENT_ID")
	clientSecret := os.Getenv("DISCORD_CLIENT_SECRET")

	requestBody := url.Values{}
	requestBody.Set("client_id", clientID)
	requestBody.Set("client_secret", clientSecret)
	requestBody.Set("grant_type", "authorization_code")
	requestBody.Set("code", code)
	requestBody.Set("redirect_uri", redirectURL)
	requestBody.Set("scope", scope)

	client := &http.Client{}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(requestBody.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	getAccessTokenResponse := &getAccessTokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(getAccessTokenResponse)

	if err != nil {
		panic(err)
	}

	return getAccessTokenResponse.AccessToken
}

// DiscordUserData contains response from Discord
type DiscordUserData struct {
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

// GetUserData gets user data from Discord
func GetUserData(accesstoken string) DiscordUserData {
	client := &http.Client{}
	req, err := http.NewRequest("GET", userDataURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accesstoken))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	discordUserData := DiscordUserData{}

	err = json.NewDecoder(resp.Body).Decode(&discordUserData)

	if err != nil {
		panic(err)
	}

	return discordUserData
}
