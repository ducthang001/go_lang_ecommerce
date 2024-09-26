package controller

import (
	"github.com/ducthang001/go-ecommerce-backend-api/internal/services"
	"github.com/ducthang001/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserById(c *gin.Context) {
	
	response.SuccessResponse(c, 20001, []string{"abc", "123", "ducthang"})
	response.ErrorResponse(c, 20003, "No need!")
}