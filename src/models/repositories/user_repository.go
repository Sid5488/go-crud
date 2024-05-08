package repositories

import (
	rest_err "github.com/Sid5488/go-crud/src/configurations"
	"github.com/Sid5488/go-crud/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(
	database *mongo.Database,
) *userRepository {
	return &userRepository{
		database,
	}
}

type UserRepository interface {
	CreateUser(
		userDomain models.UserDomainInterface,
	) (models.UserDomainInterface, *rest_err.RestErr)
}
