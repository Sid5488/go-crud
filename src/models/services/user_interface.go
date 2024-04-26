package services

import (
	rest_err "github.com/Sid5488/go-crud/src/configurations"
	"github.com/Sid5488/go-crud/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	SignUp(models.UserDomainInterface) *rest_err.RestErr

	UpdateUser(string, models.UserDomainInterface) *rest_err.RestErr

	FindUser(string) (*models.UserDomainInterface, *rest_err.RestErr)

	DeleteUser(string) *rest_err.RestErr
}
