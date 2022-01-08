package services

import (
	"context"
	"mggers-reports-api/models"
	"mggers-reports-api/utils"
)

type ReportsService interface {
	Create(ctx context.Context, request models.CreateReport) (*models.Report, error)
}

type reportsServiceImpl struct {
	repository RepositoryService
}

func InitReportsServiceImpl(config utils.AppConfig, repository RepositoryService) (*reportsServiceImpl, error) {

	return &reportsServiceImpl{repository: repository}, nil
}

func (service *reportsServiceImpl) Create(ctx context.Context, request models.CreateReport) (*models.Report, error) {
	var response models.Report

	// TODO: call repository
	// service.db.CreateReport()

	return &response, nil
}