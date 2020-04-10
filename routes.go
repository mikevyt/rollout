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
		"UsersIndex",
		"GET",
		"/user",
		h.UsersIndex,
	},
}
