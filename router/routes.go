package router

import (
	"net/http"

	handler "github.com/haffjjj/myblog-api/handlers"
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
		"GetPosts",
		"GET",
		"/posts",
		handler.GetPost,
	},
}
