package router

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/internal/handlers"
	"mggers-reports-api/internal/services"
)

func Init(service *services.Service) *gin.Engine {
	router := gin.New()
	setupRoutes(router, service)
	return router
}

func setupRoutes(r *gin.Engine, s *services.Service) {
	r.GET("/api/find", handlers.GetReportsEndpoint(s))
	r.GET("/api/insert", handlers.InsertOneReport(s))
}
