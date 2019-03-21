package models

// Post is struct for model post
type Post struct {
	Title       string   `json:"title" bson:"title" validate:"required"`
	Thumbnail   string   `json:"thumbnail" bson:"thumbnail" validate:"required"`
	CreatedAt   string   `json:"createdAt" bson:"createdAt" validate:"required"`
	ReadingTime string   `json:"readingTime" bson:"readingTime" validate:"required"`
	Tag         []string `json:"tag" bson:"tag" validate:"required"`
	Content     string   `json:"content" bson:"content" validate:"required"`
}

//PostsGroup is struct for postGroup
type PostsGroup struct {
	Count int    `json:"count" bson:"count"`
	Data  []Post `json:"data" bson:"data"`
}
