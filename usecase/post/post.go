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

func (pU *postUsecase) GetGroups(p models.Pagination) ([]*models.PostsGroup, error) {
	postsGroups, err := pU.postRepo.GetGroups(p)

	if err != nil {
		return nil, err
	}

	return postsGroups, nil
}

func (pU *postUsecase) GetGroupsByTag(t string, p models.Pagination) ([]*models.PostsGroup, error) {
	postsGroups, err := pU.postRepo.GetGroupsByTag(t, p)

	if err != nil {
		return nil, err
	}

	return postsGroups, nil
}