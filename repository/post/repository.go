package post

import "github.com/haffjjj/myblog-backend/models"

// Repository interface represent post repository contract
type Repository interface {
	GetGroups() ([]*models.PostsGroup, error)
}
