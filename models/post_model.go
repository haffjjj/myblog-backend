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

//GetPosts is model for getPosts
func GetPosts() []*Post {

	var client = mongo.MongoSession.Client
	var collection = client.Database("myblog").Collection("posts")

	var posts []*Post

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &elem)
		// fmt.Println(cur.Current)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return posts
}
