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
	log.Infof("!!!!! unregister handler,seqId=%v, opcode=%v",req.SeqID, req.Opcode())

}
func (*GwHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session) {
	log.Infof("########recv server HeartBeat,seqId=%v cmd=%v",cmd.SeqID ,cmd.Opcode())
}
func (*GwHandlers) LoginGuest(cmd *rpc.SimpleCmd, se *tcp.Session, devID string, deviceType string, partnerID string, version string){

}

func (*GwHandlers) LoginPlatform(cmd *rpc.SimpleCmd, se *tcp.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string){
	
}
