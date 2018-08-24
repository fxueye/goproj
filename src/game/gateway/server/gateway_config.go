package server

type GatewayConfig struct {
	ServerVersion string
	ServerPort int
	WsPort int
	CenterIp string
	CenterPort int
	PackLimit int
	RPCTimeOut int
	DesKey string
	SendChanLimit int
	RecvChanLimit int
}
