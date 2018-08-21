package server

type GatewayConfig struct {
	ServerVersion string
	ServerPort int
	PackLimit int
	RPCTimeOut int
	DesKey string
	SendChanLimit int
	RecvChanLimit int
}
