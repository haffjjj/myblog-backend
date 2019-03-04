package models

// Post is struct for model post
type Post struct {
	Title       string   `json:"title" bson:"title"`
	Thumbnail   string   `json:"thumbnail" bson:"thumbnail"`
	CreatedAt   string   `json:"createdAt" bson:"createdAt"`
	ReadingTime string   `json:"readingTime" bson:"readingTime"`
	Tag         []string `json:"tag" bson:"tag"`
	Content     string   `json:"content" bson:"content"`
}

//PostsGroup is struct for postGroup
type PostsGroup struct {
	Count int    `json:"count" bson:"count"`
	Data  []Post `json:"data" bson:"data"`
}
