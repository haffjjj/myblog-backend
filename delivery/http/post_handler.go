package http

import (
	"encoding/json"
	"net/http"

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

	r.HandleFunc("/postsGroups", handler.GetGroups).Methods("GET")
}

// GetGroups ...
func (p *PostHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	postsGroups := p.pUsecase.GetGroups()

	result, _ := json.Marshal(postsGroups)
	w.Write(result)
}
