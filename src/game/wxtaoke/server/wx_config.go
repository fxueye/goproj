package server

type WxConfig struct {
	QrcodeDir           string
	ControllerUserNames []string
	Groups              []string
	LoginUrl            string
	Special             []string
	EmailConfig         EmailConfig
	GroupMsg            bool
	TextConfig          []string
}
type EmailConfig struct {
	EmailAcc      string
	EmailPassword string
	SmtpServer    string
	ToEmail       []string
}
