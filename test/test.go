package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Post is struct for model post
type Post struct {
	Title       string   `json:"title" bson:"title"`
	Thumbnail   string   `json:"thumbnail" bson:"thumbnail"`
	CreatedAt   string   `json:"createdAt" bson:"createdAt"`
	ReadingTime string   `json:"readingTime" bson:"readingTime"`
	Tag         []string `json:"tag" bson:"tag"`
	Content     string   `json:"content" bson:"content"`
}

//PostPagination is struct for post pagination
type postsGroup struct {
	Count int    `json:"count" bson:"count"`
	Data  []Post `json:"data" bson:"data"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var collection = client.Database("myblog").Collection("posts")

	var postsGroups []*postsGroup

	// pipeline := mongo.Pipeline{}

	//aggregate to get data
	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"count", bson.D{{"$sum", 1}}},
				{"data", bson.D{{"$push", "$$ROOT"}}},
			}},
		},
		bson.D{
			{"$unwind", "$data"},
		},
		bson.D{
			{"$replaceRoot", bson.D{
				{"newRoot", bson.D{
					{"$mergeObjects", bson.A{"$data", "$$ROOT"}},
				}},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"count", bson.D{
					{"$first", "$count"},
				}},
				{"data", bson.D{
					{"$push", "$data"},
				}},
			}},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem postsGroup
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		postsGroups = append(postsGroups, &elem)

		// fmt.Println(cur.Current)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	res, _ := json.Marshal(postsGroups[0])

	fmt.Println(string(res))
}
