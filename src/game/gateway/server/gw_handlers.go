package server

import (
	rpc "game/common/rpc/simple"
	"game/common/server"
	wraps "game/cmds/wraps"
	log "github.com/cihub/seelog"
)

type GwHandlers struct {
	rpc.SimpleInvoker
}

func GwProxyHandler(req *rpc.SimpleCmd, se *server.Session) {
	

}
func (*GwHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *server.Session, player *wraps.PlayerWrap) {
	log.Infof("!!!!! HeartBeat handler,seqId=%v, opcode=%v ,guid=%v ,createTime=%v", cmd.SeqID, cmd.Opcode(),player.GUID,player.CreateTime)
}
func (*GwHandlers) LoginGuest(cmd *rpc.SimpleCmd, se *server.Session, devID string, deviceType string, partnerID string, version string) {

}

func (*GwHandlers) LoginPlatform(cmd *rpc.SimpleCmd, se *server.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string) {

}
