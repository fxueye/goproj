package server

import (
	conf "game/common/config"
	"game/common/server"
	"reflect"

	log "github.com/cihub/seelog"
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
	webInstance.Get("/index/(.*)", index)
	Instance.RegServ("web", webInstance)
	Instance.RegSigCallback(GWOnSignal)
}
func index(val string) string {
	var i int
	value := reflect.ValueOf(i)
	log.Infof("%v", value.Kind())
	log.Infof("%v", value.Type())
	return "hello " + val + "\n"
}
