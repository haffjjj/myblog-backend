package models

// User represent model for user data
type User struct {
	Fullname string `json:"fullname" bson:"fullname"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
