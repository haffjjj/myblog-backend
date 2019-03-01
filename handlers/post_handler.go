package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/haffjjj/myblog-api/models"
)

//GetPost is handlers
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	getPosts := models.GetPosts()[0]
	postsGroup, _ := json.Marshal(getPosts)

	// // w.Write(posts)
	// fmt.Println(posts)

	w.Write(postsGroup)
}
