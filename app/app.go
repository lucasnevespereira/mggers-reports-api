package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mggers-reports-api/router"
	"mggers-reports-api/services"
	"mggers-reports-api/utils"
)

type App struct {
	Config utils.AppConfig
	Router *gin.Engine
	Services *services.Services
}

func New() *App {
	app := &App{}
	app.setup()
	return app
}

func (a *App) setup() {
	config := utils.LoadConfig()
	services := services.Init(config)
	router := router.Init(services)

	a.Config = config
	a.Router = router
	a.Services = services
}

func (a *App) Run() {
	port := a.Config.Port
	a.Router.Run(fmt.Sprintf(":%d", port))
}
