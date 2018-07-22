package server

type WxConfig struct {
	QrcodeDir        string
	ForwardUserNames []string
	LoginUrl         string
	Special          []string
	EmailConfig      EmailConfig
	WebConfig        WebConfig
	GroupMsg         bool
	TextConfig       []string
}
type EmailConfig struct {
	EmailAcc      string
	EmailPassword string
	SmtpServer    string
	ToEmail       []string
}

type WebConfig struct {
	ServerPort int
	StaticDir  string
}
