package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mggers-reports-api/internal/models"
	"mggers-reports-api/internal/services"
	"mggers-reports-api/utils"
	"net/http"
	"time"
)

func GetReportsEndpoint(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		var results []models.Report
		find, err := service.DB.Find(c.Request.Context(), bson.D{})
		if err != nil {
			utils.Logger.Errorf("finding report: %v", find)
		}

		for find.Next(context.TODO()) {
			//Create a value into which the single document can be decoded
			var r models.Report
			err := find.Decode(&r)
			if err != nil {
				utils.Logger.Error(err)
			}

			results =append(results, r)

		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"results": results,
		})
	}
}

func InsertOneReport(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		r := models.Report{
			ID: primitive.NewObjectID(),
			Description: "test description",
			Latitude:    23,
			Longitude:   43,
			ReportedAt:  time.Now(),
		}

		one, err := service.DB.InsertOne(c.Request.Context(), r)
		if err != nil {
			utils.Logger.Errorf("inserting report: %v", one)
		}

		utils.Logger.Info("inserted report with id", one.InsertedID)

		c.JSON(http.StatusCreated, gin.H{
			"status": http.StatusCreated,
			"id": one.InsertedID,
		})
	}
}