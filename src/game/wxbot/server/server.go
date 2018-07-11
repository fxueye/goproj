package server

import (
	conf "game/common/config"
	"game/common/server"
)

type WxServer struct {
	server.Server
}

var (
	Instance   *WxServer
	config     WxConfig
	wxInstance *WxService
)

func Init() {
	Instance = &WxServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/wx_config.json", &config)
	wxInstance = newWxService(config.QrcodeDir)
	Instance.RegServ("wx", wxInstance)
	Instance.RegSigCallback(GWOnSignal)
}
