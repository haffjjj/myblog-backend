package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/haffjjj/myblog-backend/models"

	"github.com/gorilla/mux"
	"github.com/haffjjj/myblog-backend/usecase/post"
)

//PostHandler ...
type PostHandler struct {
	pUsecase post.Usecase
}

//NewPostHandler ...
func NewPostHandler(r *mux.Router, p post.Usecase) {
	handler := &PostHandler{p}

	r.HandleFunc("/postsGroups", handler.GetGroups).Queries("start", "{start}", "limit", "{limit}").Methods("GET")
	r.HandleFunc("/postsGroups", handler.GetGroups).Methods("GET")

	r.HandleFunc("/postsGroups/tag/{tag}", handler.GetGroupsByTag).Queries("start", "{start}", "limit", "{limit}").Methods("GET")
	r.HandleFunc("/postsGroups/tag/{tag}", handler.GetGroupsByTag).Methods("GET")

}

// GetGroups ...
func (p *PostHandler) GetGroups(w http.ResponseWriter, r *http.Request) {

	pagination := models.Pagination{Start: 0, Limit: 100}

	if startP, ok := mux.Vars(r)["start"]; ok {
		startP, _ := strconv.Atoi(startP)
		pagination.Start = startP
	}

	if limitP, ok := mux.Vars(r)["limit"]; ok {
		limitP, _ := strconv.Atoi(limitP)
		pagination.Limit = limitP
	}

	postsGroups, err := p.pUsecase.GetGroups(pagination)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ResponseError{Message: err.Error()})
	} else {
		json.NewEncoder(w).Encode(postsGroups)
	}
}

//GetGroupsByTag ...
func (p *PostHandler) GetGroupsByTag(w http.ResponseWriter, r *http.Request) {

	tag := mux.Vars(r)["tag"]

	pagination := models.Pagination{Start: 0, Limit: 100}

	if startP, ok := mux.Vars(r)["start"]; ok {
		startP, _ := strconv.Atoi(startP)
		pagination.Start = startP
	}

	if limitP, ok := mux.Vars(r)["limit"]; ok {
		limitP, _ := strconv.Atoi(limitP)
		pagination.Limit = limitP
	}

	postsGroups, err := p.pUsecase.GetGroupsByTag(tag, pagination)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ResponseError{Message: err.Error()})
	} else {
		json.NewEncoder(w).Encode(postsGroups)
	}
}
