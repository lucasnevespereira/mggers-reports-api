package utils

import (
	"github.com/spf13/viper"
	"strings"
)

type AppConfig struct {
	Env string
	Port int
}

func LoadConfig() AppConfig {
    viper.AutomaticEnv()
    viper.SetEnvPrefix("conf")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    viper.SetDefault("env", "dev")
    viper.SetDefault("port", 5002)

    conf := AppConfig{}
    conf.Env = viper.GetString("env")
    conf.Port = viper.GetInt("port")

    return conf
}
