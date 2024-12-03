package controller

import (
	"goGinTemplate/web/dto"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUserData() ([]dto.UserRequest, error)
}

type UserControllerImpl struct {
	service UserController
}

func UserControllerInit(userService UserController) *UserControllerImpl {
	return &UserControllerImpl{
		service: userService,
	}
}

func (srv UserControllerImpl) GetAllUserData(c *gin.Context) {
	users, err := srv.service.GetAllUserData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
