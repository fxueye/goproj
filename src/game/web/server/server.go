package server

import (
	conf "game/common/config"
	"game/common/server"
	// log "github.com/cihub/seelog"
)

type WebServer struct {
	server.Server
}

var (
	Instance    *WebServer
	config      WebConfig
	webInstance *WebService
)

func Init() {
	Instance = &WebServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/web_config.json", &config)
	webInstance = newWebService(config.ServerPort, config.StaticDir)
	Instance.RegServ("web", webInstance)
	Instance.RegSigCallback(OnSignal)
}
