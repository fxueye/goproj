package server

import (
	log "github.com/cihub/seelog"
	"time"
	"sync"
	"sync/atomic"
	"game/common/server/tcp"
	rpc "game/common/rpc/simple"
	cmd "game/cmds"
	// log "github.com/cihub/seelog"
)

type CS2GWService struct{
	tcp.ISessionHandler
	*tcp.TcpService
	simpleRPC *rpc.SimpleRPC
	sessionBySID map[int64] *tcp.Session
	sessLock sync.RWMutex
	sessionId int64
}
func newCS2GWService(port int) *CS2GWService{
	serv := new(CS2GWService)
	inv := cmd.NewServerCSCmdsInvoker(&GWHandlers{},ProxyHandler)
	serv.simpleRPC = rpc.NewSimpleRPC(inv,true,time.Duration(config.RPCTimeOut) * time.Second,nil)
	serv.TcpService = tcp.NewTcpService(
		port,
		time.Second,
		serv.simpleRPC,
		serv,
		tcp.SessionConfig{config.SendChanLimit,config.RecvChanLimit})
	serv.sessionBySID= make(map[int64]*tcp.Session)
	return serv
}
func (serv *CS2GWService) OnConnect(se *tcp.Session) bool{
	log.Debugf("on connnected ,addr = %v",se.GetConn().RemoteAddr().String())
	seId:= atomic.AddInt64(&serv.sessionId,1)
	se.Sid = seId
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	serv.sessionBySID[seId] = se
	return true
}
func (serv *CS2GWService) OnClose(se *tcp.Session){
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	delete(serv.sessionBySID,se.Sid)
}
func (serv *CS2GWService) OnMessage(se *tcp.Session,p tcp.IPacket) bool{
	defer func(){
		if err := recover(); err != nil{
			log.Error(err)
			se.Close()
		}
	}()
	serv.simpleRPC.Process(se,p)
	return true
}