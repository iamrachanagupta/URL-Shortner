package initapp

import (
	"urlapp/internal/utilities"

	log "github.com/sirupsen/logrus"
)

//Initialize function reads the configuration file and initializes all the application parameters
//It initializes the log format and log level as per the config
func Initialize() {
	utilities.ReadConfig()
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	log.SetLevel(log.Level(utilities.Data["appconfig"].Loglevel))
}
