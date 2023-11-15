package repository

import "fitness-app/internal/app/domain"

// UserRepository defines the interface for user repository
type UserRepository interface {
	GetByID(id int) (*domain.User, error)
	Create(user *domain.User) error
	// Add more methods as needed
}
