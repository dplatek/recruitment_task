package logger

import (
	"log"
	"os"
)

func SetLogLevel(level string) {
	switch level {
	case "Debug":
		log.SetOutput(os.Stdout)
	case "Info":
		log.SetOutput(os.Stdout)
	case "Error":
		log.SetOutput(os.Stderr)
	default:
		log.Printf("Unknown log level: %s. Defaulting to Info.", level)
		log.SetOutput(os.Stdout)
	}
}
