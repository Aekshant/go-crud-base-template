package controller

import (
	"net/http"
	"strconv"
	dto "test/internal/domain/dto"
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", uc.GetUser)
		userGroup.POST("/", uc.CreateUser)
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	// Parse the user ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the service to retrieve the user
	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Respond with the user data
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// Parse the request body into a DTO

	var userDto dto.CreateUserRequest
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call the service to create a new user
	user, err := uc.userService.CreateUser(&userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with the created user data
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
