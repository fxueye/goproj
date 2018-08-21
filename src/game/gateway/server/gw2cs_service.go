package server

import (
	"sync"
	"game/common/server/tcp"
	rpc "game/common/rpc/simple"
)

type GW2GSService struct{
	tcp.ISessionHandler
	*tcp.TcpService
	simpleRPC *rpc.SimpleRPC
	sessionBySID map[int64] *tcp.Session
	sessLock sync.RWMutex
}