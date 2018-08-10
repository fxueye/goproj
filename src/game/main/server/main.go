package main

import (
	"os"

	log "github.com/cihub/seelog"
)

var (
	prof interface {
		Stop()
	}
	TimeDelta int64 //本地时间与ntp差值
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
	GetNtpTime()
	log.Info("gateway server closed")
}
