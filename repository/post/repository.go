package post

import "github.com/haffjjj/myblog-backend/models"

// Repository interface represent post repository contract
type Repository interface {
	GetGroups(p models.Pagination) ([]*models.PostsGroup, error)
	GetGroupsByTag(t string, p models.Pagination) ([]*models.PostsGroup, error)
}
