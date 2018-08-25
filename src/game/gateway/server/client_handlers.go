package server

import (
	wraps "game/cmds/wraps"
	rpc "game/common/rpc/simple"
	"game/common/server"
)

type ClientHandlers struct {
	rpc.SimpleInvoker
}

func ClientProxyHandler(cmd *rpc.SimpleCmd, se *server.Session) {
	if cmd.Opcode() < 10000 { //转发至客户端

	} else if cmd.Opcode() < 20000 { //转发至cs

	}

}
func (*ClientHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *server.Session) {

}
func (*ClientHandlers) LoginSuccess(cmd *rpc.SimpleCmd, se *server.Session, player *wraps.PlayerWrap, reconnect bool, extension string) {

}
func (*ClientHandlers) LoginFailed(cmd *rpc.SimpleCmd, se *server.Session, errorCode int16, errMsg string) {

}
