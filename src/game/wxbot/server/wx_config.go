package server

type WxConfig struct {
	QrcodeDir        string
	ForwardUserNames []string
	LoginUrl         string
	Special          []string
}
