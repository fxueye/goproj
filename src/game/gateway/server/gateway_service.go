package server

import (
	log "github.com/cihub/seelog"
	"sync/atomic"
	"time"
	"sync"
	cmd "game/cmds"
	rpc "game/common/rpc/simple"
	tcp "game/common/server/tcp"
)
const (
	CHECK_INTERNAL int = 5
)
type GatewayService struct{
	tcp.ISessionHandler
	*tcp.TcpService
	simpleRPC  *rpc.SimpleRPC
	sessions map[int64]*tcp.Session
	sessLock sync.RWMutex
	sessionId int64
	pkgLimit int
}
func newGatewayService(port,pkgLimit int) *GatewayService{
	serv := new(GatewayService)
	inv := cmd.NewClientCmdsInvoker(&ClientHandlers{},ClientProxyHandler)
	serv.simpleRPC = rpc.NewSimpleRPC(inv,false,time.Duration(config.RPCTimeOut) * time.Second,[]byte(config.DesKey))
	serv.TcpService = tcp.NewTcpService(
		port,
		time.Second,
		serv.simpleRPC,
		serv,
		tcp.SessionConfig{config.SendChanLimit,config.RecvChanLimit})
	serv.sessions = make(map[int64] *tcp.Session)
	serv.pkgLimit = pkgLimit
	return serv
}
func (serv *GatewayService) OnConnect(se *tcp.Session) bool{
	log.Debugf("on connected, addr = %v",se.GetConn().RemoteAddr())
	seId := atomic.AddInt64(&serv.sessionId,1)
	se.Sid = seId
	
	serv.AsyncDo(func(){
		for{
			if se.IsClosed(){
				return
			}
			se.SetAttr("pkgcnt",int(0))
			if heartbeatArr, ok := se.GetAttr("heartbeat");ok{
				if hbtime,yes := heartbeatArr.(time.Time); yes{
					if time.Now().After(hbtime.Add(time.Second * 120)){
						se.Close()
						return
					}
				}
			}else{

			}
			time.Sleep(time.Second * time.Duration(CHECK_INTERNAL))
		}
	})
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	serv.sessions[seId] = se
	log.Debugf("sessionId %d",seId)
	return true
}
func (serv *GatewayService) OnClose(se *tcp.Session){
	log.Debugf("on close, addr = %v sid[%v]",se.GetConn().RemoteAddr(),se.Sid)
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	delete(serv.sessions,se.Sid)
}
func (serv *GatewayService) OnMessage(se *tcp.Session,p tcp.IPacket) bool{
	defer func(){
		if err := recover(); err != nil{
			log.Error(err)
			se.Close()
		}
	}()
	if cmd,ok := p.(*rpc.SimpleCmd); ok{
		cmd.SetSID(se.Sid)
		se.SetAttr("heartbeat",time.Now());
		if serv.pkgLimit > 0 {
			if pkgcnt,y := se.GetAttr("pkgcnt");y{
				if pkgcnt.(int) >  serv.pkgLimit * CHECK_INTERNAL{
					log.Errorf("To many packet addr[%v] pkgcnt[%v]",se.GetConn().RemoteAddr(),pkgcnt.(int))
					se.Close()
					return false
				}
				se.SetAttr("pkgcnt",pkgcnt.(int) + 1)
			}else{
				se.SetAttr("pkgcnt",1)
			}
		}
		serv.simpleRPC.Process(se,p)
	}else{
		log.Debug("not simplecmd Message!")
	}
	return true
}

func (serv *GatewayService) Boradcast(opID int16 , data []byte){
	serv.sessLock.RLock()
	defer serv.sessLock.RUnlock()
	for _,v := range serv.sessions{
		serv.simpleRPC.Send(v,int16(0),opID,int64(0),data)
	}
}