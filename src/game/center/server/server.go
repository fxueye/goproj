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
)

func Init() {
	Instance = &CenterServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/center_config.json", &config)
	Instance.RegSigCallback(OnSignal)
}
