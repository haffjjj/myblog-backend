package post

import (
	"github.com/haffjjj/myblog-backend/models"
	"github.com/haffjjj/myblog-backend/repository/post"
)

type postUsecase struct {
	postRepo post.Repository
}

//NewPostUsecase ...
func NewPostUsecase(p post.Repository) Usecase {
	return &postUsecase{
		postRepo: p,
	}
}

func (p *postUsecase) GetGroups() ([]*models.PostsGroup, error) {
	postsGroups, err := p.postRepo.GetGroups()

	if err != nil {
		return nil, err
	}

	return postsGroups, nil
}
