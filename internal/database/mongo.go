package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mggers-reports-api/utils"
)

func NewMongoClient(conf utils.MongoConfig) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, err
	}

	return client, nil
}
