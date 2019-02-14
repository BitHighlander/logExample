package main
import (
	"github.com/joho/godotenv"
	"os"
)
import log "github.com/sirupsen/logrus"

func main() {
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Get Log level from config
	var logLevel = os.Getenv("DEFAULT_LOG_LEVEL")
	if logLevel == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	log.Debug("This is only shown on dev")
	log.Info("wuddup")
	log.Warn("watch out bro")
	log.Error("Shits fucked")
}
