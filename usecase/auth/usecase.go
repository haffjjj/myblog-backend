package auth

import (
	"github.com/haffjjj/myblog-backend/models"
)

//Usecase represent auth usecase contract
type Usecase interface {
	Auth(u, p string) (*models.User, error)
}
