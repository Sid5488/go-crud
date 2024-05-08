package repositories

import (
	"context"
	"os"

	rest_err "github.com/Sid5488/go-crud/src/configurations"
	"github.com/Sid5488/go-crud/src/configurations/logger"
	"github.com/Sid5488/go-crud/src/models"
)

var (
	MONGODB_USER_DATABASE = "MONGODB_USER_DATABASE"
)

func (ur *userRepository) CreateUser(
	userDomain models.UserDomainInterface,
) (models.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	colletion_name := os.Getenv(MONGODB_USER_DATABASE)
	collection := ur.databaseConnection.Collection(colletion_name)

	value, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil
}
