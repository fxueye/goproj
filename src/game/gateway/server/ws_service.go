package server

import (
	cmd "game/cmds"
	"game/common/server/ws"
	"time"
	"game/common/server"
	rpc "game/common/rpc/simple"
	log "github.com/cihub/seelog"
	"sync"
	"golang.org/x/net/websocket"
	"sync/atomic"
)

type WsService struct {
	server.ISessionHandler
	*ws.WebsocketService
	simpleRPC *rpc.SimpleRPC
	sessions  map[int64]*server.Session
	sessLock  sync.RWMutex
	sessionId int64
	pkgLimit  int
}

func newWsService(port int, pkgLimit int) *WsService {
	serv := new(WsService)
	inv := cmd.NewClientCmdsInvoker(&ClientHandlers{}, ClientProxyHandler)
	serv.simpleRPC = rpc.NewSimpleRPC(inv, false, time.Duration(config.RPCTimeOut)*time.Second, nil)
	serv.WebsocketService = ws.NewWebsocketService(
		port, 
		time.Second,
		serv.simpleRPC,
		serv,
		server.SessionConfig{config.SendChanLimit, config.RecvChanLimit})
	
	serv.sessions = make(map[int64]*server.Session)
	serv.pkgLimit = pkgLimit
	return serv
}
func (serv *WsService) OnConnect(se *server.Session) bool {
	log.Debugf("on connected, addr = %v", se.GetConn().(*websocket.Conn).RemoteAddr())
	seId := atomic.AddInt64(&serv.sessionId, 1)
	se.Sid = seId

	// serv.AsyncDo(func() {
	// 	for {
	// 		if se.IsClosed() {
	// 			return
	// 		}
	// 		se.SetAttr("pkgcnt", int(0))
	// 		if heartbeatArr, ok := se.GetAttr("heartbeat"); ok {
	// 			if hbtime, yes := heartbeatArr.(time.Time); yes {
	// 				if time.Now().After(hbtime.Add(time.Second * 120)) {
	// 					se.Close()
	// 					return
	// 				}
	// 			}
	// 		} else {

	// 		}
	// 		time.Sleep(time.Second * time.Duration(CHECK_INTERNAL))
	// 	}
	// })
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	serv.sessions[seId] = se
	log.Debugf("sessionId %d", seId)
	return true
}
func (serv *WsService) OnClose(se *server.Session) {
	log.Debugf("on close, addr = %v sid[%v]", se.GetConn().(*websocket.Conn).RemoteAddr(), se.Sid)
	serv.sessLock.Lock()
	defer serv.sessLock.Unlock()
	delete(serv.sessions, se.Sid)
}
func (serv *WsService) OnMessage(se *server.Session, p server.IPacket) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			se.Close()
		}
	}()
	if cmd, ok := p.(*rpc.SimpleCmd); ok {
		cmd.SetSID(se.Sid)
		se.SetAttr("heartbeat", time.Now())
		if serv.pkgLimit > 0 {
			if pkgcnt, y := se.GetAttr("pkgcnt"); y {
				if pkgcnt.(int) > serv.pkgLimit*CHECK_INTERNAL {
					log.Errorf("To many packet addr[%v] pkgcnt[%v]", se.GetConn().(*websocket.Conn).RemoteAddr(), pkgcnt.(int))
					se.Close()
					return false
				}
				se.SetAttr("pkgcnt", pkgcnt.(int)+1)
			} else {
				se.SetAttr("pkgcnt", 1)
			}
		}
		serv.simpleRPC.Process(se, p)
	} else {
		log.Debug("not simplecmd Message!")
	}
	return true
}

func (serv *WsService) Boradcast(opID int16, data []byte) {
	serv.sessLock.RLock()
	defer serv.sessLock.RUnlock()
	for _, v := range serv.sessions {
		serv.simpleRPC.Send(v, int16(0), opID, int64(0), data)
	}
}
