package main

import (
	"log"
	"net/http"

	"github.com/haffjjj/myblog-api/db/mongo"
	"github.com/haffjjj/myblog-api/router"
	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

func init() {
	mongo.Connect()
}

func main() {

	router := router.NewRouter()

	err := http.ListenAndServe(":8000", setupGlobalMiddleware(router))

	if err == nil {
		log.Fatal(err)
	}

}
