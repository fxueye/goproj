package server

import (
	cmds "game/cmds"
	rpc "game/common/rpc/simple"
	"game/common/server"

	log "github.com/cihub/seelog"
)

type GwHandlers struct {
	rpc.SimpleInvoker
}

func ProxyHandler(cmd *rpc.SimpleCmd, se *server.Session) {
	log.Infof("########recv ProxyHandler ,seqId=%v, cmd=%v", cmd.SeqID, cmd.Opcode())
	op := cmd.Opcode()
	if op < 10000 {

	} else if op > 20000 && op < 21000 { //转发给gateway

	} else {
		log.Errorf("center not register[%d]", cmd.Opcode())
	}
}
func (*GwHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *server.Session) {
	log.Infof("########recv client HeartBeat,seqId=%v, cmd=%v", cmd.SeqID, cmd.Opcode())
	cs2GwInstance.simpleRPC.Send(se, cmd.SeqID(), cmds.ServerGWCmds_HEART_BEAT, 0)
}
func (*GwHandlers) LoginGuest(cmd *rpc.SimpleCmd, se *server.Session, devID string, deviceType string, partnerID string, version string) {

}

func (*GwHandlers) LoginPlatform(cmd *rpc.SimpleCmd, se *server.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string) {

}
