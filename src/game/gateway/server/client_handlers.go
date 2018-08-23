package server
import(
	rpc "game/common/rpc/simple"
	tcp "game/common/server/tcp"
	wraps "game/cmds/wraps"
)
type ClientHandlers struct {
	rpc.SimpleInvoker
}

func ClientProxyHandler(cmd *rpc.SimpleCmd, se *tcp.Session) {
	if cmd.Opcode() < 10000 { //转发至客户端
		
	} else if cmd.Opcode() < 20000 { //转发至cs

	}

}
func(*ClientHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session){

}
func(*ClientHandlers) LoginSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, player *wraps.PlayerWrap, reconnect bool, extension string){

}
func(*ClientHandlers) LoginFailed(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int16, errMsg string){

}
