package controllers

import (
	"github.com/Sid5488/go-crud/src/models/services"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface services.UserDomainService,
) *userControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	SignUp(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}
