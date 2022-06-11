package utils

import (
	"os"
)

type AppConfig struct {
	AppName string
	Version string
	Env     string
	Port    int
	Mongo   MongoConfig
}

type MongoConfig struct {
	URI        string
	Database   string
	Collection string
}

func LoadConfig() AppConfig {
	conf := AppConfig{}

	conf.AppName = "mggers-reports-api"
	conf.Version = "0.0"
	conf.Env = os.Getenv("ENV")
	conf.Port = 9000
	conf.Mongo.URI = os.Getenv("MONGO_SERVER")
	conf.Mongo.Database = "mggers-database"
	reportsCol := "reports"
	if conf.Env == "dev" {
		reportsCol = "reports-dev"
	}
	conf.Mongo.Collection = reportsCol

	return conf
}
