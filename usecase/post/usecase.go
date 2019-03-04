package post

import "github.com/haffjjj/myblog-backend/models"

//Usecase interface represent post usecase contract
type Usecase interface {
	GetGroups() []*models.PostsGroup
}
