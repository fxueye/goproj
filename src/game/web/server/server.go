package server

import (
	conf "game/common/config"
	"game/common/server"
	"game/common/utils"
	"io/ioutil"
	"runtime"

	log "github.com/cihub/seelog"
)

type WebServer struct {
	server.Server
}

var (
	Instance       *WebServer
	config         WebConfig
	webInstance    *WebService
	webTlsInstance *WebTlsService
)

func Init() {
	Instance = &WebServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/web_config.json", &config)
	pkey, _ := ioutil.ReadFile(config.Tls.Pkey)
	cert, _ := ioutil.ReadFile(config.Tls.Cert)
	log.Infof("pkey:\n%v", string(pkey))
	log.Infof("cert:\n%v", string(cert))
	webTlsInstance = newWebTlsService(config.ServerTlsPort, config.StaticTlsDir, cert, pkey)
	webInstance = newWebService(config.ServerPort, config.StaticDir)
	Instance.RegServ("web", webInstance)
	Instance.RegServ("webTls", webTlsInstance)
	Instance.RegSigCallback(OnSignal)
}

func ShowStack() {
	buf := make([]byte, 1<<20)
	runtime.Stack(buf, false)
	log.Error("============Panic Stack Info===============")
	log.Errorf("\n%s", buf)

	if config.EmailConfig.EmailAcc != "" {
		str := string(buf)
		for _, email := range config.EmailConfig.ToEmail {
			err := utils.SendMail(config.EmailConfig.EmailAcc, config.EmailConfig.EmailPassword, config.EmailConfig.SmtpServer, email, "server crash", str, "")
			if err != nil {
				log.Error(err)
			}
		}
	}
}
