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
	app.Config = conf

	client, err := database.Connect(context.Background(), conf.Mongo)
	if err != nil {
		return err
	}

	db := database.New(client, conf.Mongo)
	utils.Logger.Info("db -> %v", db)

	service := services.New(db, conf.Mongo)
	r := router.New(service)

	app.Router = r

	return nil
}

func (app *App) Run() {
	port := app.Config.Port
	utils.Logger.Infof("Running app on port %d", port)
	app.Router.Run(fmt.Sprintf(":%d", port))
}
