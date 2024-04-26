package views

import (
	"github.com/Sid5488/go-crud/src/controllers/models/response"
	"github.com/Sid5488/go-crud/src/models"
)

func ConvertDomainToResponse(
	userDomain models.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		Id:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
