package repository

import (
	"errors"
	dao "test/internal/domain/dao"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *dao.User) error
	GetUserByID(ID uint) (*dao.User, error)
	// ... other methods
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Implement methods for CreateUser, GetUserByID, etc.
func (repo *userRepository) CreateUser(user *dao.User) error {
	// Use GORM to save the user to the database
	result := repo.db.Create(user)
	return result.Error
}

func (repo *userRepository) GetUserByID(ID uint) (*dao.User, error) {
	var user dao.User
	result := repo.db.First(&user, ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, result.Error // Other errors 1
	}
	return &user, nil
}
