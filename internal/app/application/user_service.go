package application

import (
	"fitness-app/internal/app/domain"
	"fitness-app/internal/app/repository"
)

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

// GetUserByID retrieves a user by ID
func (s *UserServiceImpl) GetUserByID(id int) (*domain.User, error) {
	return s.UserRepository.GetByID(id)
}

// AddUser adds a new user
func (s *UserServiceImpl) AddUser(user *domain.User) error {
	return s.UserRepository.Create(user)
}
