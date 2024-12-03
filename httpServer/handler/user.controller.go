package handler

import (
	dto "goGinTemplate/domain/dto"
	"log"
)

type UserController interface {
	GetAllUserData()
}

type UserService interface {
	GetAllUser() ([]dto.GetUserDto, error)
}

type UserControllerImpl struct {
	service UserService
}

func UserControllerInit(userService UserService) *UserControllerImpl {
	return &UserControllerImpl{
		service: userService,
	}
}

func (srv *UserControllerImpl) GetAllUserData() {
	_, err := srv.service.GetAllUser()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, common.ServiceRepo[[]dto.GetUserDto]{
		// 	Data:    nil,
		// 	Err:     err,
		// 	Success: false,
		// })
		return
	}
	log.Println("Success-------------")
	// c.JSON(http.StatusOK, common.ServiceRepo[[]dto.GetUserDto]{
	// 	Data:    data,
	// 	Err:     nil,
	// 	Success: true,
	// })
}
