package server

type CenterConfig struct {
	ServerVersion string
	ServerPort int
	PackLimit int
	RPCTimeOut int
	DesKey string
	SendChanLimit int
	RecvChanLimit int
}
