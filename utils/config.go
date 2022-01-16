package utils

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppName string
	Version string
	Env  string
	Port int
	Mongo MongoConfig
}

type MongoConfig struct {
	URI string
	Database string
	Collection string
}

func LoadConfig() AppConfig {
	viper.AutomaticEnv()
	viper.SetDefault("app_name", "mggers-reports-api")
	viper.SetDefault("version", "0.0")
	viper.SetDefault("env", "dev")
	viper.SetDefault("port", 9000)

	viper.SetDefault("mongo_server", "mongodb://mongodb")
	viper.SetDefault("mongo_database", "mggers-reports")
	viper.SetDefault("mongo_collection", "Reports")

	conf := AppConfig{}

	conf.Env = viper.GetString("app_name")
	conf.Env = viper.GetString("version")
	conf.Env = viper.GetString("env")
	conf.Port = viper.GetInt("port")

	conf.Mongo.URI = viper.GetString("mongo_server")
	conf.Mongo.Database = viper.GetString("mongo_database")
	conf.Mongo.Collection = viper.GetString("mongo_collection")

	return conf
}
