package server

type WxConfig struct {
	QrcodeDir        string
	ForwardUserNames []string
	LoginUrl         string
	Special          []string
	DBConfig         DBConfig
	EmailConfig      EmailConfig
}
type DBConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBMaxOpen  int
	DBMaxIdle  int
}
type EmailConfig struct {
	EmailAcc      string
	EmailPassword string
	SmtpServer    string
	ToEmail       []string
}
