package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"mggers-reports-api/database"
	"mggers-reports-api/router"
	"mggers-reports-api/services"
	"mggers-reports-api/utils"
)

type App struct {
	Router *gin.Engine
	Config utils.AppConfig
}

func New() *App {
	app := &App{}
	if err := app.setup(); err != nil {
		utils.Logger.Errorf("app.setup: %v", err)
	}
	return app
}

func (app *App) setup() error {
	utils.Logger.Info("setting up app")

	conf := utils.LoadConfig()
	client, err := database.Connect(context.Background(), conf.Mongo)
	if err != nil {
		return err
	}

	db := database.New(client, conf.Mongo.Database)

	service := services.New(db)
	r := router.New(service)

	app.Router = r
	app.Config = conf

	return nil
}

func (app *App) Run() {
	port := app.Config.Port
	utils.Logger.Infof("Running %s on port %d", app.Config.AppName, port)
	app.Router.Run(fmt.Sprintf(":%d", port))
}
