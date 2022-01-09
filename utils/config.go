package utils

import (
	"github.com/spf13/viper"
	"strings"
)

type AppConfig struct {
	Env             string
	Port            int
	PostgreHost     string
	PostgrePort     int
	PostgreUser     string
	PostgrePassword string
	PostgreDbName   string
	PostgreSsl      string
}

func LoadConfig() AppConfig {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("conf")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("env", "dev")
	viper.SetDefault("port", 5002)

	viper.SetDefault("pg.host", "localhost")
	viper.SetDefault("pg.port", 5432)
	viper.SetDefault("pg.username", "__")
	viper.SetDefault("pg.password", "__")
	viper.SetDefault("pg.database", "__")
	viper.SetDefault("pg.ssl", "disable")

	conf := AppConfig{}
	conf.Env = viper.GetString("env")
	conf.Port = viper.GetInt("port")

	conf.PostgreHost = viper.GetString("pg.host")
	conf.PostgrePort = viper.GetInt("pg.port")
	conf.PostgreUser = viper.GetString("pg.username")
	conf.PostgrePassword = viper.GetString("pg.password")
	conf.PostgreDbName = viper.GetString("pg.database")
	conf.PostgreSsl = viper.GetString("pg.ssl")

	return conf
}
