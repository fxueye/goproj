package server

import (
	conf "game/common/config"
	"game/common/server"
	"game/common/utils"
	"game/wxbot/db"
	"runtime"

	log "github.com/cihub/seelog"
)

type WxServer struct {
	server.Server
}

const WXSERVICE = "wx"
const WEBSERVICE = "web"

var (
	Instance    *WxServer
	config      WxConfig
	wxInstance  *WxService
	webInstance *WebService
	DBMgr       *db.DBMgr
)

func Init() {
	Instance = &WxServer{
		server.NewServer(),
	}
	conf.LoadConfig("json", "config/wx_config.json", &config)
	DBMgr = new(db.DBMgr)
	DBMgr.Init()
	err := DBMgr.CreateWxDB(config.DBConfig.DBHost, config.DBConfig.DBPort, config.DBConfig.DBUser, config.DBConfig.DBPassword, config.DBConfig.DBName, config.DBConfig.DBMaxOpen, config.DBConfig.DBMaxIdle)
	if err != nil {
		log.Error(err)
		return
	}
	wxInstance = newWxService(config.LoginUrl, config.QrcodeDir)
	webInstance = newWebService(config.WebConfig.ServerPort, config.WebConfig.StaticDir)

	Instance.RegServ(WXSERVICE, wxInstance)
	Instance.RegServ(WEBSERVICE, webInstance)

	Instance.RegSigCallback(GWOnSignal)
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
