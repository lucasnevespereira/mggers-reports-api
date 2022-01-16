package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mggers-reports-api/database"
	"mggers-reports-api/router"
	"mggers-reports-api/services"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	app := &App{}
	if err := app.setup(); err != nil {
		log.Printf("setting up: %v", err)
	}
	return app
}

func (app *App) setup() error {
	fmt.Println("setting up app")

	client, err := database.Init(context.Background())
	if err != nil {
		return err
	}

	db := database.New(client, "mggers-reports")

	service := services.New(db)
	r := router.New(service)
	app.router = r

	return nil
}


func (app *App) Run() {
	port := 9000
	app.router.Run(fmt.Sprintf(":%d", port))
}
