package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mggers-reports-api/models"
	"mggers-reports-api/utils"
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

	one, err := reportsCollection.InsertOne(ctx, models.Report{
		Description: "single desc",
		Latitude:    34,
		Longitude:   123,
		ReportedAt:  time.Now(),
	})
	if err != nil {
		utils.Logger.Errorf("inserting one: %v", err)
	}

	fmt.Println(one.InsertedID)

}
