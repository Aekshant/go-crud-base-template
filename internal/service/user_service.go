package service

import (
	dao "test/internal/domain/dao"
	dto "test/internal/domain/dto"
	"test/internal/repository"
)

type UserService interface {
	CreateUser(userDto *dto.CreateUserRequest) (*dao.User, error)
	GetUserByID(ID uint) (*dao.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(userDto *dto.CreateUserRequest) (*dao.User, error) {
	user := &dao.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password, // Consider hashing the password
	}

	err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(ID uint) (*dao.User, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
