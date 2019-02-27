package models

import (
	"context"
	"fmt"
	"log"

	"github.com/haffjjj/myblog-api/db/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Post is struct for model post
type Post struct {
	Title       string   `json:"title"`
	Thumbnail   string   `json:"thumbnail"`
	Tag         []string `json:"tag"`
	CreatedAt   string   `json:"createdAt"`
	ReadingTime string   `json:"readingTime"`
	Content     string   `json:"content"`
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
		fmt.Println(cur.Current)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return posts
}
