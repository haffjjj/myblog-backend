package main

import (
	"context"
	"log"

	_postRepo "github.com/haffjjj/myblog-backend/repository/post"
	_postUsecase "github.com/haffjjj/myblog-backend/usecase/post"
	"github.com/labstack/echo"

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

	e := echo.New()
	_httpDelivery.NewPostHandler(e, postUsecase)

	e.Logger.Fatal(e.Start(":8081"))
}
