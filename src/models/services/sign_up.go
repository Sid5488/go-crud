package services

import (
	"fmt"

	rest_err "github.com/Sid5488/go-crud/src/configurations"
	"github.com/Sid5488/go-crud/src/configurations/logger"
	"github.com/Sid5488/go-crud/src/models"
	"go.uber.org/zap"
)

func (ud *userDomainService) SignUp(userDomain models.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(ud)

	return nil
}
