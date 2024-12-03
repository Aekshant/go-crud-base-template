package handler

import (
	"goGinTemplate/app/common"
	dto "goGinTemplate/domain/dto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUserData(c *gin.Context)
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

func (srv *UserControllerImpl) GetAllUserData(c *gin.Context) {
	data, err := srv.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ServiceRepo[[]dto.GetUserDto]{
			Data:    nil,
			Err:     err,
			Success: false,
		})
		return
	}
	log.Println("Success-------------")
	c.JSON(http.StatusOK, common.ServiceRepo[[]dto.GetUserDto]{
		Data:    data,
		Err:     nil,
		Success: true,
	})
}
