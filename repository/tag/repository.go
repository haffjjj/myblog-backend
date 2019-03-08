package tag

import "github.com/haffjjj/myblog-backend/models"

// Repository ...
type Repository interface {
	Get() ([]*models.Tag, error)
}
