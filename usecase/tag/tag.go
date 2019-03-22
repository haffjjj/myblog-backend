package tag

import (
	"github.com/haffjjj/myblog-backend/models"
	"github.com/haffjjj/myblog-backend/repository/tag"
)

type tagUsecase struct {
	tagRepo tag.Repository
}

//NewTagUsecase represent initialitation tagusecase
func NewTagUsecase(t tag.Repository) Usecase {
	return &tagUsecase{
		tagRepo: t,
	}
}

func (t *tagUsecase) Get() ([]*models.Tag, error) {
	tags, err := t.tagRepo.Get()

	if err != nil {
		return nil, err
	}

	return tags, nil
}
