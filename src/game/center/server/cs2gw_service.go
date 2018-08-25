package server

import (
	cmd "game/cmds"
	rpc "game/common/rpc/simple"
	"game/common/server"
	"game/common/server/tcp"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/cihub/seelog"
	// log "github.com/cihub/seelog"
)

type CS2GWService struct {
	server.ISessionHandler
	*tcp.TcpService
	simpleRPC    *rpc.SimpleRPC
	sessionBySID map[int64]*server.Session
	sessLock     sync.RWMutex
	sessionId    int64
}

func newCS2GWService(port int) *CS2GWService {
	serv := new(CS2GWService)
	inv := cmd.NewServerGWCmdsInvoker(&GwHandlers{}, ProxyHandler)
	serv.simpleRPC = rpc.NewSimpleRPC(inv, true, time.Duration(config.RPCTimeOut)*time.Second, nil)
	serv.TcpService = tcp.NewTcpService(
		port,
		time.Second,
		serv.simpleRPC,
		serv,
		server.SessionConfig{config.SendChanLimit, config.RecvChanLimit})
	serv.sessionBySID = make(map[int64]*server.Session)
	return serv
}
func (serv *CS2GWService) OnConnect(se *server.Session) bool {
	log.Debugf("on connnected ,addr = %v", se.GetConn().RemoteAddr().String())
	seId := atomic.AddInt64(&serv.sessionId, 1)
	se.Sid = seId
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	serv.sessionBySID[seId] = se
	return true
}
func (serv *CS2GWService) OnClose(se *server.Session) {
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	delete(serv.sessionBySID, se.Sid)
}
func (serv *CS2GWService) OnMessage(se *server.Session, p server.IPacket) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			se.Close()
		}
	}()
	serv.simpleRPC.Process(se, p)
	return true
}
