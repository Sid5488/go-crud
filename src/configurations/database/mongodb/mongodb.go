package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL           = "MONGODB_URL"
	MONGODB_USER_DATABASE = "MONGODB_USER_DATABASE"
)

func NewMongoDBConnection() (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_user_database := os.Getenv(MONGODB_USER_DATABASE)

	ctx := context.Background()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri),
	)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_user_database), nil
}
