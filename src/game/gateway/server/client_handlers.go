package server

import (
	cmds "game/cmds"
	wraps "game/cmds/wraps"
	rpc "game/common/rpc/simple"
	"game/common/server"
	log "github.com/cihub/seelog"
	"time"
)

type ClientHandlers struct {
	rpc.SimpleInvoker
}

func ClientProxyHandler(cmd *rpc.SimpleCmd, se *server.Session) {
	if cmd.Opcode() < 10000 { //转发至客户端

	} else if cmd.Opcode() < 20000 { //转发至cs

	}

}

func (*ClientHandlers) HeartBeat(cmd *rpc.SimpleCmd, se *server.Session, player *wraps.PlayerWrap) {
	log.Infof("!!!!! HeartBeat handler,seqId=%v, opcode=%v ,guid=%v ,createTime=%v", cmd.SeqID, cmd.Opcode(),player.GUID,player.CreateTime)
	wsInstance.simpleRPC.Send(se, 0, cmds.ClientCmds_HEART_BEAT, 0,"你好!")
	var p = new(wraps.PlayerWrap)
	p.GUID = "10001";
	p.CreateTime = time.Now().Unix()
	wsInstance.simpleRPC.Send(se,0,cmds.ClientCmds_LOGIN_SUCCESS,0,p,false,"")
}
func (*ClientHandlers) LoginGuest(cmd *rpc.SimpleCmd, se *server.Session, devID string, deviceType string, partnerID string, version string) {

}
func (*ClientHandlers) LoginPlatform(cmd *rpc.SimpleCmd, se *server.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string) {

}
