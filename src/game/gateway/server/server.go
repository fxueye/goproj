package server

import (
	conf "game/common/config"
	"game/common/server"
)

type GateWayServer struct {
	server.Server
}

var (
	Instance   *GateWayServer
	config     GateConfig
	gsInstance *GateService
)

func Init() {
	Instance = &GateWayServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/gate_config.json", &config)
	gsInstance = newGateService(config.ServerPort)
	Instance.RegServ("gs", gsInstance)
	Instance.RegSigCallback(GWOnSignal)
}
