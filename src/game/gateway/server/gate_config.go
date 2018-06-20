package server

type GateConfig struct {
	ServerVersion string
	ServerPort    int
	SendChanLimit int
	RecvChanLimit int
}
