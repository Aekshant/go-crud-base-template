package usersusecases

import (
	dao "goGinTemplate/domain/dao"
	dto "goGinTemplate/domain/dto"
	"log"
)

type UserService interface {
	GetAllUser() ([]dto.GetUserDto, error)
}

type UserRepository interface {
	FindAllUser() ([]dao.User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func UserServiceInit(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u UserServiceImpl) GetAllUser() ([]dto.GetUserDto, error) {
	data, err := u.userRepository.FindAllUser()
	if err != nil {
		return nil, err
	}

	var UserResponses []dto.GetUserDto
	for _, user := range data {
		log.Println(user.Name)
		log.Println(user.Email)
		UserResponses = append(UserResponses, dto.GetUserDto{
			ID:    uint(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return UserResponses, nil
}
