package router

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/handlers"
)

func Init() *gin.Engine {
	r := gin.New()
	setupRoutes(r)
	return r
}

func setupRoutes(r *gin.Engine) {
	r.GET("/", handlers.GetHealth)
}