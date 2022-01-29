package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mggers-reports-api/internal/models"
	"mggers-reports-api/utils"
)

type ReportService interface {
	Create(request models.ReportRequest) (*models.Report, error)
	GetAll(ctx context.Context) (*[]models.Report, error)
}

func (s *Service) Create(ctx context.Context, request models.ReportRequest) (*models.Report, error) {

	r := models.Report{
		ID:          primitive.NewObjectID(),
		Description: request.Description,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		ReportedAt:  request.ReportedAt,
	}

	one, err := s.DB.InsertOne(ctx, r)
	if err != nil {
		utils.Logger.Errorf("inserting report: %v", one)
		return nil, err
	}

	utils.Logger.Info("inserted report with ", one.InsertedID)

	return &r, nil
}

func (s *Service) GetAll(ctx context.Context) (*[]models.Report, error) {

	var results []models.Report
	find, err := s.DB.Find(ctx, bson.D{})
	if err != nil {
		utils.Logger.Errorf("finding report: %v", find)
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
