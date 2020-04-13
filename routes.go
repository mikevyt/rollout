package main

import (
	"net/http"

	h "github.com/mikevyt/rollout/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"DiscordRedirect",
		"GET",
		"/",
		h.DiscordRedirect,
	},
	Route{
		"Login",
		"GET",
		"/login",
		h.Login,
	},
	Route{
		"GetUsers",
		"GET",
		"/user",
		h.GetUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/user/{discordid}",
		h.GetUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{discordid}",
		h.DeleteUser,
	},
}
