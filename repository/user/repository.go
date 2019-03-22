package user

import (
	"github.com/haffjjj/myblog-backend/models"
)

//Repository represent user repository contract
type Repository interface {
	GetByUsername(u string) (*models.User, error)
}
