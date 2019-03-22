package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Post represent model for post data
type Post struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Thumbnail   string             `json:"thumbnail" bson:"thumbnail" validate:"required"`
	CreatedAt   string             `json:"createdAt" bson:"createdAt" validate:"required"`
	ReadingTime string             `json:"readingTime" bson:"readingTime" validate:"required"`
	Tag         []string           `json:"tag" bson:"tag" validate:"required"`
	Content     string             `json:"content" bson:"content" validate:"required"`
}

//PostsGroup represent for postGroup data
type PostsGroup struct {
	Count int    `json:"count" bson:"count"`
	Data  []Post `json:"data" bson:"data"`
}
