package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haffjjj/myblog-api-backend/db/mongo"
	"github.com/haffjjj/myblog-api-backend/router"
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

	fmt.Println("Starting sever..")
	router := router.NewRouter()
	err := http.ListenAndServe(":8000", setupGlobalMiddleware(router))
	if err == nil {
		log.Fatal(err)
	}

}
