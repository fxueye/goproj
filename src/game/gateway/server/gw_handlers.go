package server

import (
	rpc "game/common/rpc/simple"
	"game/common/server"

	log "github.com/cihub/seelog"
)

type GwHandlers struct {
	rpc.SimpleInvoker
}

func GwProxyHandler(req *rpc.SimpleCmd, se *server.Session) {
	log.Infof("!!!!! unregister handler,seqId=%v, opcode=%v", req.SeqID, req.Opcode())

}
func (*GwHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *server.Session) {
	// log.Infof("########recv server HeartBeat,seqId=%v cmd=%v", cmd.SeqID, cmd.Opcode())
}
func (*GwHandlers) LoginGuest(cmd *rpc.SimpleCmd, se *server.Session, devID string, deviceType string, partnerID string, version string) {

}

func (*GwHandlers) LoginPlatform(cmd *rpc.SimpleCmd, se *server.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string) {

}
