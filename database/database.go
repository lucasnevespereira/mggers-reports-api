package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mggers-reports-api/utils"
)

func Connect(ctx context.Context, conf utils.MongoConfig) (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			utils.Logger.Errorf("disconnect: %v", err)
		}
	}(client, ctx)

	return client, nil
}


func New(client *mongo.Client, conf utils.MongoConfig) *mongo.Database {
	return client.Database(conf.Database)
}
