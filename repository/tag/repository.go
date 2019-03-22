package tag

import "github.com/haffjjj/myblog-backend/models"

// Repository represent tag repository contract
type Repository interface {
	Get() ([]*models.Tag, error)
}
