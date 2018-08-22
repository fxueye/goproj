package server


import(
	"time"
	log "github.com/cihub/seelog"
	simple "game/common/rpc/simple"
	tcp "game/common/server/tcp"
)
type GWHandlers struct {
	simple.SimpleInvoker
}

func ProxyHandler(cmd *simple.SimpleCmd, se *tcp.Session) {
	log.Infof("########recv ProxyHandler , cmd=%v", cmd.Opcode())
	op := cmd.Opcode()
	if op < 10000 {

	} else if op > 20000 && op < 21000 { //转发给gateway

	} else {
		log.Errorf("charaterServer not register[%d]", cmd.Opcode())
	}
}
func (*GWHandlers) GW2CS_Ping(cmd *simple.SimpleCmd, se *tcp.Session) {
	log.Infof("########recv client HeartBeat, cmd=%v", cmd.Opcode())
	cs2GwInstance.simpleRPC.Send(se,cmd.SeqID(),0,0,time.Now().UnixNano());
}
func (*GWHandlers) GW2CS_LoginGuest(cmd *simple.SimpleCmd, se *tcp.Session, deviceID string, deviceType string, partnerID string, gameversion string) {

}