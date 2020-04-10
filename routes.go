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
		"PostUser",
		"POST",
		"/user",
		h.PostUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{discordid}",
		h.DeleteUser,
	},
}
