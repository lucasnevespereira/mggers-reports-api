package services

import (
	"gorm.io/gorm"
	"mggers-reports-api/utils"
)

type RepositoryService interface {

}

type RepositoryServiceImpl struct {
	db *gorm.DB
}

func NewRepository(config utils.AppConfig) (*RepositoryServiceImpl, error) {
	// TODO: Open DB with Gorm
	// db, err := gorm.Open()

	return &RepositoryServiceImpl{db: nil}, nil
}