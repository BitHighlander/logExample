package main

import (
	"github.com/joho/godotenv"
	"os"
)
import log "github.com/sirupsen/logrus"

func main() {
	//load env
	godotenv.Load()

	//Get Log level from config
	var logLevel = os.Getenv("DEFAULT_LOG_LEVEL")
	if logLevel == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else {
		// NOTE: (if no file assume production)
		log.SetLevel(log.InfoLevel)
	}

	log.Debug("This is only shown on dev")
	log.Info("wuddup")
	log.Warn("watch out bro")
	log.Error("Shits fucked")
}
