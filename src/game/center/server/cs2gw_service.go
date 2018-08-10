package server

import (
	"sync"
	"game/common/server/tcp"
	rpc "game/common/rpc/simple"
)

type CS2GWService struct{
	tcp.ISessionHandler
	*tcp.TcpService
	simpleRPC *rpc.SimpleRPC
	sessionBySID map[int64] *tcp.Session
	sessLock sync.RWMutex
}