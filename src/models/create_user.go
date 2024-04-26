package models

import (
	"fmt"

	rest_err "github.com/Sid5488/go-crud/src/configurations"
	"github.com/Sid5488/go-crud/src/configurations/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
