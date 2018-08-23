package cmds
import (
	"errors"
	"fmt"
	rpc "game/common/rpc/simple"
	tcp "game/common/server/tcp"
	
)

type IServerGWCmds interface {
	HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session) // 心跳
	LoginGuest(cmd *rpc.SimpleCmd, se *tcp.Session, devID string, deviceType string, partnerID string, version string) // 登录
	LoginPlatform(cmd *rpc.SimpleCmd, se *tcp.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string) // 登录(ptID平台）
	
}

type ServerGWCmdsInvoker struct {
	invoker IServerGWCmds
	defaultInvoker func(cmd *rpc.SimpleCmd, se *tcp.Session)
	rpc.SimpleInvoker
} 

func NewServerGWCmdsInvoker(invoker IServerGWCmds, defaultInvoker func(*rpc.SimpleCmd, *tcp.Session)) *ServerGWCmdsInvoker {
	inv := new(ServerGWCmdsInvoker)
	inv.invoker = invoker
	inv.defaultInvoker = defaultInvoker
	return inv
} 

func (this *ServerGWCmdsInvoker) Invoke(cmd *rpc.SimpleCmd, se *tcp.Session) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	pack := cmd.Pack()
	switch(cmd.Opcode()) {
	case 0: 
		this.invoker.HeartBeat(cmd,se)
	case 10001: 
		this.invoker.LoginGuest(cmd,se, pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString())
	case 10002: 
		this.invoker.LoginPlatform(cmd,se, pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopBool(), pack.PopString(), pack.PopString())
	
	default:
		if this.defaultInvoker != nil {
			this.defaultInvoker(cmd,se)
		}
	}
	return nil
}

