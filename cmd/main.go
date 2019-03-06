package main

import (
	"context"
	"fmt"
	"log"

	_postRepo "github.com/haffjjj/myblog-backend/repository/post"
	_postUsecase "github.com/haffjjj/myblog-backend/usecase/post"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_httpDelivery "github.com/haffjjj/myblog-backend/delivery/http"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func init() {
	viper.SetConfigFile("config")
	viper.SetConfigType("json")
	viper.SetConfigFile("config/config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var (
		dbHost = viper.GetString("database.mongodb.host")
		dbPort = viper.GetString("database.mongodb.port")
		port   = viper.GetString("port")
	)

	mgoClient, err := mongo.Connect(context.TODO(), fmt.Sprint("mongodb://", dbHost, dbPort))
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

	e.Logger.Fatal(e.Start(port))
}
