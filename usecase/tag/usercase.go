package tag

import "github.com/haffjjj/myblog-backend/models"

//Usecase ...
type Usecase interface {
	Get() ([]*models.Tag, error)
}
