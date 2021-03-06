package main

import (
	"net/http"

	"github.com/mikevyt/rollout/handlers"
)

// Route contains the Name, Method, Pattern and Handler for a specific endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a slice containing Route
type Routes []Route

var routes = Routes{
	Route{
		"DiscordCallback",
		"GET",
		"/auth/{provider}/callback",
		handlers.DiscordCallback,
	},
	Route{
		"Logout",
		"GET",
		"/logout/{provider}",
		handlers.Logout,
	},
	Route{
		"Authenticate",
		"GET",
		"/auth/{provider}",
		handlers.Authenticate,
	},
	Route{
		"GetUsers",
		"GET",
		"/user",
		handlers.GetUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/user/{discordid}",
		handlers.GetUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{discordid}",
		handlers.DeleteUser,
	},
}
