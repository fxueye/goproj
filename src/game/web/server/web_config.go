package server

type WebConfig struct {
	ServerVersion string
	ServerPort    int
	StaticDir     string

	ServerTlsPort int
	StaticTlsDir  string

	WeiAppid    string
	WeiSecret   string
	WeiApiUrl   string
	Tls         TLS
	EmailConfig EmailConfig
}
type EmailConfig struct {
	EmailAcc      string
	EmailPassword string
	SmtpServer    string
	ToEmail       []string
}
type TLS struct {
	Pkey string
	Cert string
}
