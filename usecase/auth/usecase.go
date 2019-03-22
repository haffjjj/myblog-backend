package auth

import (
	"github.com/haffjjj/myblog-backend/models"
)

//Usecase ...
type Usecase interface {
	Auth(u, p string) (*models.User, error)
}
