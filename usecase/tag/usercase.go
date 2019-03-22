package tag

import "github.com/haffjjj/myblog-backend/models"

//Usecase represent tag usecase contract
type Usecase interface {
	Get() ([]*models.Tag, error)
}
