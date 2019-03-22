package post

import "github.com/haffjjj/myblog-backend/models"

//Usecase represent post usecase contract
type Usecase interface {
	GetByID(i string) (*models.Post, error)
	GetGroups(p models.Pagination) ([]*models.PostsGroup, error)
	GetGroupsByTag(t string, p models.Pagination) ([]*models.PostsGroup, error)
	Store(p *models.Post) error
	Delete(i string) error
	Update(i string, p *models.Post) error
}
