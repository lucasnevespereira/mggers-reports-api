package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DB *mongo.Database
}

func New(db *mongo.Database) *Service {
	return &Service{DB: db}
}
