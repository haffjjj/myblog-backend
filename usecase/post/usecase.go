package post

import "github.com/haffjjj/myblog-backend/models"

//Usecase interface represent post usecase contract
type Usecase interface {
	GetGroups(p models.Pagination) ([]*models.PostsGroup, error)
	GetGroupsByTag(t string, p models.Pagination) ([]*models.PostsGroup, error)
}
