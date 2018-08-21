package server
import(
	simple "game/common/rpc/simple"
	tcp "game/common/server/tcp"
)
type ClientHandlers struct {
	simple.SimpleInvoker
}

func ProxyHandler(cmd *simple.SimpleCmd, se *tcp.Session) {
	if cmd.Opcode() < 10000 { //转发至客户端
		
	} else if cmd.Opcode() < 20000 { //转发至cs

	}

}
func (*ClientHandlers) HeartBeat(cmd *simple.SimpleCmd, se *tcp.Session) {

}
func (*ClientHandlers) LoginGuest(cmd *simple.SimpleCmd, se *tcp.Session, deviceID string, deviceType string, partnerID string, gameversion string) {

}