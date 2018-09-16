package server

import (
	"fmt"
	conf "game/common/config"
	"game/common/server"
	"game/common/utils"
	"net/http"
	"runtime"

	log "github.com/cihub/seelog"
)

type WxServer struct {
	server.Server
}

const WXSERVICE = "wx"

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
	wxInstance = newWxService(config.LoginUrl, config.QrcodeDir, config.TempImgDir)
	go WebService(config.QrcodeDir)
	Instance.RegServ(WXSERVICE, wxInstance)
	Instance.RegSigCallback(OnSignal)
}
func WebService(path string) {
	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Infof("server start on:%d", config.WebPort)
	http.ListenAndServe(fmt.Sprintf(":%d", config.WebPort), nil)
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
