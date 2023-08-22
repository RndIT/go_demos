package main

import (
	"os"

	utils "github.com/RndIT/go_demos/internal/Utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)

}

func main() {
	// инициализация контекста
	ctxLogger := log.WithFields(log.Fields{
		"module": utils.GetFunctionName(main),
		"other":  "I also should be logged always",
	})

	ctxLogger.Info("I'll be logged with common and other field")
	ctxLogger.Info("Me too")
}
