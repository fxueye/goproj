package server

import (
	conf "game/common/config"
	"game/common/server"

	// log "github.com/cihub/seelog"
)

type GatewayServer struct {
	server.Server
}

var (
	Instance    *GatewayServer
	config GatewayConfig
	gwInstance *GatewayService
)

func Init() {
	Instance = &GatewayServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/gateway_config.json", &config)
	gwInstance = newGatewayService(config.ServerPort,config.PackLimit)
	Instance.RegServ("gw",gwInstance)
	Instance.RegSigCallback(OnSignal)
}
