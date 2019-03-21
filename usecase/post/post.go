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

func (pU *postUsecase) Update(i string, p *models.Post) error {
	err := pU.postRepo.Update(i, p)
	if err != nil {
		return err
	}

	return nil
}

//Delete ...
func (pU *postUsecase) Delete(i string) error {
	err := pU.postRepo.Delete(i)
	if err != nil {
		return err
	}

	return nil
}

func (pU *postUsecase) Store(p *models.Post) error {
	err := pU.postRepo.Store(p)
	if err != nil {
		return err
	}

	return nil
}

//GetById ...
func (pU *postUsecase) GetByID(i string) (*models.Post, error) {
	post, err := pU.postRepo.GetByID(i)

	if err != nil {
		return nil, err
	}

	return post, nil
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
