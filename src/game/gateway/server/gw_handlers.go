package server
import (
	log "github.com/cihub/seelog"
	rpc "game/common/rpc/simple"
	tcp "game/common/server/tcp"
)
type GwHandlers struct {
	rpc.SimpleInvoker
}
func GwProxyHandler(req *rpc.SimpleCmd, se *tcp.Session) {
	log.Infof("!!!!! unregister handler, opcode=%v", req.Opcode())

}
func (*GwHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session) {
	log.Infof("########recv client HeartBeat, cmd=%v", cmd.Opcode())
}
func (*GwHandlers) GW2CS_Ping(cmd *rpc.SimpleCmd, se *tcp.Session){

}
func (*GwHandlers) GW2CS_LoginGuest(cmd *rpc.SimpleCmd, se *tcp.Session, deviceID string, deviceType string, partnerID string, ip string) {

}