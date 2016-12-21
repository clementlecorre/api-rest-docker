package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"deploy",
		"POST",
		"/deploy",
		deploy,
	},
	// Route{
	// 	"deploytest",
	// 	"POST",
	// 	"/deploytest",
	// 	deploytest,
	// },
	// Route{
	// 	"TodoCreate",
	// 	"POST",
	// 	"/todos",
	// 	TodoCreate,
	// },
	// Route{
	// 	"TodoShow",
	// 	"GET",
	// 	"/todos/{todoId}",
	// 	TodoShow,
	// },
}
