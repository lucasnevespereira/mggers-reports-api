package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mggers-reports-api/router"
	"mggers-reports-api/utils"
)

type App struct {
	Config utils.AppConfig
	Router *gin.Engine
}

func New() *App {
	app := &App{}
	app.setup()
	return app
}

func (a *App) setup() {
	a.Config = utils.LoadConfig()
	a.Router = router.Init()
}

func (a *App) Run() {
	port := a.Config.Port
	a.Router.Run(fmt.Sprintf(":%d", port))
}
