package server

import (
	conf "game/common/config"
	"game/common/server"

	// log "github.com/cihub/seelog"
)

type CenterServer struct {
	server.Server
}

var (
	Instance    *CenterServer
	config CenterConfig
	cs2GwInstance *CS2GWService;
)

func Init() {
	Instance = &CenterServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/center_config.json", &config)
	cs2GwInstance = newCS2GWService(config.ServerPort)
	Instance.RegServ("cs2gw",cs2GwInstance)
	Instance.RegSigCallback(OnSignal)
}
