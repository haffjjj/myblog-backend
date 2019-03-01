package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/haffjjj/myblog-api-backend/models"
)

//GetPostsGroups is handlers
func GetPostsGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	filter := models.GetPostsGroupsFilter{
		Tag: "",
	}

	getPostsGroups := models.GetPostsGroups(filter)
	postsGroups, _ := json.Marshal(getPostsGroups)

	w.Write(postsGroups)
}

//GetPostsGroupsTag is handlers
func GetPostsGroupsTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	filter := models.GetPostsGroupsFilter{
		Tag: mux.Vars(r)["tag"],
	}

	getPostsGroups := models.GetPostsGroups(filter)
	postsGroups, _ := json.Marshal(getPostsGroups)

	w.Write(postsGroups)
}
