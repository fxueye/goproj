package server

type WxConfig struct {
	QrcodeDir           string
	TempImgDir          string
	ControllerUserNames []string
	Groups              []string
	SendTimer           []string
	LoginUrl            string
	Special             []string
	EmailConfig         EmailConfig
	TextConfig          []string
	AppKey              string
	AppSecret           string
	AppPid              string
	KeyWords            []string
}
type EmailConfig struct {
	EmailAcc      string
	EmailPassword string
	SmtpServer    string
	ToEmail       []string
}
