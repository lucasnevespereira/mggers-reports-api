package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DB *mongo.Collection
}

func New(db *mongo.Collection) *Service {
	return &Service{
		DB: db,
	}
}
