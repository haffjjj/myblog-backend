package main

import (
	"context"
	"log"
	"net/http"

	_postRepo "github.com/haffjjj/myblog-backend/repository/post"
	_postUsecase "github.com/haffjjj/myblog-backend/usecase/post"

	"github.com/gorilla/mux"
	_httpDelivery "github.com/haffjjj/myblog-backend/delivery/http"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {

	mgoClient, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = mgoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mgoClient.Disconnect(context.TODO())

	postRepo := _postRepo.NewMongoPostRespository(mgoClient)
	postUsecase := _postUsecase.NewPostUsecase(postRepo)

	r := mux.NewRouter()
	_httpDelivery.NewPostHandler(r, postUsecase)

	log.Fatal(http.ListenAndServe(":8081", r))
}
