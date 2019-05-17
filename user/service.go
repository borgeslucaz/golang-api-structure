package user

import (
	"github.com/borgeslucaz/golang-api-structure/models"
	"github.com/borgeslucaz/golang-api-structure/utils"
	"errors"
)

// UserService for manage all business rules
var UserService *Service

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Find a user
func (s *Service) Find(id int) (*models.User, error) {
	return s.repo.Find(id)
}
// FindByEmail find user by email
func (s *Service) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

// Create Persist uses on db
func (s *Service) Create(user *models.User) (*models.User, error) {
	return s.repo.Create(user)
}

// Login validation
func (s *Service) Login(email string, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	var hash utils.Hash
	if hash.Compare(user.Password, password) == nil {
		return user, nil
	}
	return user, errors.New("Password not match")
}
