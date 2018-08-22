package main

import (
	"os"

	log "github.com/cihub/seelog"
)

var (
	prof interface {
		Stop()
	}
)

func main() {
	defer func() {

		if err := recover(); err != nil {
			log.Critical(err)
			os.Exit(0)
		}
	}()

	logger, err := log.LoggerFromConfigAsFile("config/log.xml")
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
	log.Info("gateway server closed")
}
