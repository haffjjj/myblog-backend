package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/middleware"

	_postRepo "github.com/haffjjj/myblog-backend/repository/post"
	_tagRepo "github.com/haffjjj/myblog-backend/repository/tag"
	_userRepo "github.com/haffjjj/myblog-backend/repository/user"
	_authUsecase "github.com/haffjjj/myblog-backend/usecase/auth"
	_postUsecase "github.com/haffjjj/myblog-backend/usecase/post"
	_tagUsecase "github.com/haffjjj/myblog-backend/usecase/tag"

	"github.com/haffjjj/myblog-backend/utils"

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
		dbHost  = viper.GetString("database.mongodb.host")
		dbPort  = viper.GetString("database.mongodb.port")
		dbUName = viper.GetString("database.mongodb.username")
		dbPass  = viper.GetString("database.mongodb.password")
		port    = viper.GetString("port")
	)

	mgoClient, err := mongo.Connect(context.TODO(), fmt.Sprint("mongodb://", dbUName, ":", dbPass, "@", dbHost, dbPort))
	if err != nil {
		log.Fatal(err)
	}
	err = mgoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mgoClient.Disconnect(context.TODO())

	// ===========

	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = &utils.Validator{Validator: validator.New()}

	// ===========

	postRepo := _postRepo.NewMongoPostRespository(mgoClient)
	tagRepo := _tagRepo.NewMongoTagRespository(mgoClient)
	userRepo := _userRepo.NewMongoUserRespository(mgoClient)

	postUsecase := _postUsecase.NewPostUsecase(postRepo)
	tagUsecase := _tagUsecase.NewTagUsecase(tagRepo)
	authUsecase := _authUsecase.NewAuthUsecase(userRepo)

	_httpDelivery.NewAuthHandler(e, authUsecase)
	_httpDelivery.NewPostHandler(e, postUsecase)
	_httpDelivery.NewTagHandler(e, tagUsecase)

	// ===========

	e.Logger.Fatal(e.Start(port))
}
