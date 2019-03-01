package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/haffjjj/myblog-api-backend/models"
)

//GetPost is handlers
func GetPostsGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	getPostsGroups := models.GetPostsGroups()
	postsGroups, _ := json.Marshal(getPostsGroups)

	w.Write(postsGroups)
}
