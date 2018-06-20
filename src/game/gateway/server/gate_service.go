package server

import (
	"time"

	"game/common/server/tcp"
)

type GateService struct {
	tcp.ISessionHandler
	*tcp.TcpService
	proxies      map[int32]*tcp.Session
	sessionBySID map[int64]*tcp.Session
}


func newGateService(port int) *GateService {
	serv := new(GateService)
	serv.TcpService = tcp.NewTcpService(
		port,
		time.Second,
		,
		serv,
		tcp.SessionConfig{config.SendChanLimit, config.RecvChanLimit})
	serv.proxies = make(map[int32]*tcp.Session)
	serv.sessionBySID = make(map[int64]*tcp.Session)
	return serv
}
