package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mggers-reports-api/utils"
)

type ReportService struct {
	DB *mongo.Database
	Conf utils.MongoConfig
}

func New(db *mongo.Database, conf utils.MongoConfig) *ReportService {
	return &ReportService{DB: db, Conf: conf}
}
