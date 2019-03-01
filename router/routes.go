package router

import (
	"net/http"

	handler "github.com/haffjjj/myblog-api-backend/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetPostsGroups",
		"GET",
		"/postsGroups",
		handler.GetPostsGroups,
	},
	// Route{
	// 	"Get",
	// 	"GET",
	// 	"/posts",
	// 	handler.GetPostGroups,
	// },
}
