package server

import (
	"time"
	"sync"
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
func newCS2GWService(port int){
	serv := new(CS2GWService)
	inv := cmd.NewServerCSCmdsInvoker(&GWHandlers{},ProxyHandler)
	serv.simpleRPC = rpc.NewSimpleRPC(inv,true,time.Duration(config.RPCTimeOut) * time.Second,nil)
}