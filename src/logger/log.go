package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Logger *log.Logger

func Inst() *log.Logger {
	return Logger
}

func init() {
	Logger = log.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(log.InfoLevel)
	Logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
