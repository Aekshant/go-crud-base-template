package usersusecases

import (
	common "goGinTemplate/app/common"
	domain "goGinTemplate/domain"
	repository "goGinTemplate/infra/mysqlRepo"
)

type UserService interface {
	GetAllUser()
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
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
