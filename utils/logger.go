package utils

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var Logger = &log.Logger{
	Out: os.Stderr,
	Level: log.DebugLevel,
	Formatter: &prefixed.TextFormatter{
		DisableColors: false,
		DisableTimestamp: false,
	},
}


