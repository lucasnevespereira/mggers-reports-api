package app

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"mggers-reports-api/internal/database"
	"mggers-reports-api/internal/router"
	"mggers-reports-api/internal/services"
	"mggers-reports-api/utils"
	"net/http"
	"time"
)

type App struct {
	Config utils.AppConfig
}

func New() *App {
	return &App{
		Config: utils.LoadConfig(),
	}
}

func (app *App) Run() error {
	utils.Logger.Info("Setting up app")
	mongo, err := database.NewMongoClient(app.Config.Mongo)
	if err != nil {
		return errors.Wrap(err, "database.NewMongoClient")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = mongo.Connect(ctx)
	if err != nil {
		utils.Logger.Errorf("mongo client connect: %v", err)
	}
	defer mongo.Disconnect(ctx)

	reportsCol := mongo.Database(app.Config.Mongo.Database).
		Collection(app.Config.Mongo.Collection)

	reportService := services.New(reportsCol)

	appRouter := router.Init(reportService)

	utils.Logger.Info(fmt.Sprintf("Running %v on port %d", app.Config.AppName, app.Config.Port))

	err = http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), appRouter)
	if err != nil {
		return err
	}

	return nil
}
