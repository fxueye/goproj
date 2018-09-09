package server

import (
	conf "game/common/config"
	"game/common/server"
	"game/common/utils"
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

	Instance.RegServ(WXSERVICE, wxInstance)
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
