package services

import (
	"mggers-reports-api/utils"
)

type Services struct {
	ReportsService ReportsService
}

func Init(config utils.AppConfig) *Services {
	s := &Services{}

	repository, err := NewRepository(config)
	if err != nil {
		utils.Logger.Error(err)
	}

	reportsService, err := InitReportsServiceImpl(config, repository)
	if err != nil {
		utils.Logger.Error(err)
	}

	s.ReportsService = reportsService


	return s
}
