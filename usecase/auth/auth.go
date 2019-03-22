package auth

import (
	"errors"

	"github.com/haffjjj/myblog-backend/models"
	"github.com/haffjjj/myblog-backend/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo user.Repository
}

//NewAuthUsecase represent initialitation authusecase
func NewAuthUsecase(u user.Repository) Usecase {
	return &authUsecase{
		userRepo: u,
	}
}

func (aU *authUsecase) Auth(u, p string) (*models.User, error) {

	user, err := aU.userRepo.GetByUsername(u)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
