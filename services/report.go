package services

import (
	"context"
	"fmt"
	"mggers-reports-api/models"
	"mggers-reports-api/utils"
	"time"
)

type IReportService interface {
	CreateReport(ctx context.Context, request models.ReportRequest) (*models.Report, error)
}


func (s *ReportService) CreateReport(ctx context.Context, request models.ReportRequest) (*models.Report, error) {
	err := s.DB.Client().Connect(ctx)
	if err != nil {
		utils.Logger.Errorf("connect error: %v", err)
		return nil, err
	}
	fmt.Println("got here: %v", s.DB.Client())
	r := models.Report{
		Description: request.Description,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		ReportedAt:  time.Now(),
	}

	fmt.Printf("report: %v", r)

	one, err := s.DB.Collection(s.Conf.Collection).InsertOne(ctx,r)
	if err != nil {
		return nil, err
	}

	utils.Logger.Infof("inserted report with id %d", one.InsertedID)

	return &r, nil
}

