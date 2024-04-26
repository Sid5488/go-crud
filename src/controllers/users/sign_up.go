package controller

import (
	"net/http"

	"github.com/Sid5488/go-crud/src/configurations/logger"
	"github.com/Sid5488/go-crud/src/configurations/validation"
	"github.com/Sid5488/go-crud/src/controllers/models/request"
	"github.com/Sid5488/go-crud/src/models"
	"github.com/Sid5488/go-crud/src/models/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface services.UserDomainService
)

func SignUp(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "sign-up"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
	}

	domain := models.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := services.NewUserDomainService()
	if err := service.SignUp(domain); err != nil {
		c.JSON(err.Code, err)

		return
	}

	logger.Info("User created successfuly", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, domain)
}
