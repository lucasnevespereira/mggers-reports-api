package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mggers-reports-api/internal/models"
	"mggers-reports-api/utils"
)

type ReportService interface {
	Create(request models.ReportRequest) (*models.Report, error)
	GetAll(ctx context.Context) (*[]models.Report, error)
	DeleteById(ctx context.Context, reportId string) error
}

func (s *Service) Create(ctx context.Context, request models.ReportRequest) (*models.Report, error) {

	r := models.Report{
		ID:          primitive.NewObjectID(),
		Description: request.Description,
		Position:    request.Position,
		ReportedAt:  request.ReportedAt,
	}

	// TTL index
	index := mongo.IndexModel{
		Keys:    bson.M{"reportedAt": 1},
		Options: options.Index().SetExpireAfterSeconds(5),
	}

	s.DB.Indexes().CreateOne(ctx, index)

	one, err := s.DB.InsertOne(ctx, r)

	if err != nil {
		utils.Logger.Errorf("inserting report: %v", err)
		return nil, err
	}

	utils.Logger.Info("inserted report with ", one.InsertedID)

	return &r, nil
}

func (s *Service) GetAll(ctx context.Context) (*[]models.Report, error) {

	var results []models.Report
	find, err := s.DB.Find(ctx, bson.D{})
	if err != nil {
		utils.Logger.Errorf("finding report: %v", err)
		return nil, err
	}

	for find.Next(ctx) {
		var r models.Report
		err := find.Decode(&r)
		if err != nil {
			utils.Logger.Error(err)
		}
		results = append(results, r)
	}

	return &results, nil
}

func (s *Service) Delete(ctx context.Context, reportId string) error {
	_, err := s.DB.DeleteOne(ctx, reportId)
	if err != nil {
		utils.Logger.Errorf("db delete one: %v", err)
		return err
	}

	return nil

}
