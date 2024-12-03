package usersusecases

import (
	common "goGinTemplate/app/common"
	domain "goGinTemplate/domain"
)

type UserService interface {
	GetAllUser()
}

type userRepository interface {
	FindAllUser() ([]domain.User, error)
}

type UserServiceImpl struct {
	userRepository userRepository
}

func UserServiceInit(userRepository userRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u UserServiceImpl) GetAllUser() *common.ServiceRepo[[]domain.User] {
	// defer pkg.PanicHandler(c)

	// log.Info("start to execute get all data user")

	data, err := u.userRepository.FindAllUser()
	if err != nil {
		// log.Error("Happened Error when find all user data. Error: ", err)
		// pkg.PanicException(constant.UnknownError)
	}

	return (&common.ServiceRepo[[]domain.User]{Data: data, Err: nil, Success: true})
}
