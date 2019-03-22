package user

import (
	"github.com/haffjjj/myblog-backend/models"
)

//Repository ...
type Repository interface {
	GetByUsername(u string) (*models.User, error)
}
