package server

import (
	conf "game/common/config"
	"game/common/server"
	"runtime"
	log "github.com/cihub/seelog"
)

type GatewayServer struct {
	server.Server
}

var (
	Instance    *GatewayServer
	config GatewayConfig
	gwInstance *GatewayService
	gw2csInstance *GW2GSService
)

func Init() {
	Instance = &GatewayServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/gateway_config.json", &config)
	gwInstance = newGatewayService(config.ServerPort,config.PackLimit)
	Instance.RegServ("gw",gwInstance)
	gw2csInstance = newGW2GSService("gw2cs",config.CenterIp,config.CenterPort)
	Instance.RegServ("gw2cs",gw2csInstance)
	Instance.RegSigCallback(OnSignal)
}

func ShowStack() {
	buf := make([]byte, 1<<20)
	runtime.Stack(buf, false)
	log.Error("============Panic Stack Info===============")
	log.Errorf("\n%s", buf)
}
