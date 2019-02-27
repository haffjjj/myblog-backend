package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

//Session is struct for mongo.Client
type Session struct {
	Client *mongo.Client
}

//MongoSession is current mongo client
var MongoSession Session

//Connect will created session connection for mongo
func Connect() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	(&MongoSession).Client = client
}
