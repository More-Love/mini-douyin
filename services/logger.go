package services

import (
	"log"
	"mini-douyin/config"
	"os"
)

var logger *log.Logger

func init() {
	var ouput *os.File

	if config.Config.ServiceLoggerOutput == "" {
		ouput = os.Stdout
	} else {
		ouput, _ = os.OpenFile(config.Config.ServiceLoggerOutput, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}

	logger = log.New(ouput, "", log.Ldate|log.Ltime|log.Lshortfile)
}
