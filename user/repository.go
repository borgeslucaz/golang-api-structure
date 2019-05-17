package user

import (
	"github.com/borgeslucaz/golang-api-structure/models"
)

// Repository for user
type Repository interface {
	Find(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

