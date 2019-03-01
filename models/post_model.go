package models

import (
	"context"
	"log"

	"github.com/haffjjj/myblog-api/db/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
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

//GetPosts is model for getPosts
func GetPostsGroups() []*postsGroup {

	var client = mongo.MongoSession.Client
	var collection = client.Database("myblog").Collection("posts")

	var postsGroups []*postsGroup

	//aggregate to get data
	cur, err := collection.Aggregate(context.TODO(), []bson.D{
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
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return postsGroups
}
