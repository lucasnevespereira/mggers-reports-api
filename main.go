package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mggers-reports-api/handlers"
	"mggers-reports-api/utils"
	"net/http"
	"time"
)


func main() {

	conf := utils.LoadConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.Mongo.URI))
	if err != nil {
		utils.Logger.Errorf("mongo new client: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		utils.Logger.Errorf("mongo client connect: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(conf.Mongo.Database)
	reportsCollection := db.Collection(conf.Mongo.Collection)



	router := gin.New()
	router.GET("/find", handlers.GetReportsEndpoint(reportsCollection))
	router.GET("/insert", handlers.InsertOneReport(reportsCollection))

	http.ListenAndServe(":9000", router)

}
