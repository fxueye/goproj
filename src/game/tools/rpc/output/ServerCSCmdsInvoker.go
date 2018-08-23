package cmds
import (
	"errors"
	"fmt"
	rpc "tipcat.com/common/rpc/simple"
	tcp "tipcat.com/common/server/tcp"
	
)

type IServerCSCmds interface {
	GW2CS_Ping(cmd *rpc.SimpleCmd, se *tcp.Session) // gw心跳
	GW2CS_LoginGuest(cmd *rpc.SimpleCmd, se *tcp.Session, deviceID string, deviceType string, partnerID string, ip string) // 登录
	
}

type ServerCSCmdsInvoker struct {
	invoker IServerCSCmds
	defaultInvoker func(cmd *rpc.SimpleCmd, se *tcp.Session)
	rpc.SimpleInvoker
} 

func NewServerCSCmdsInvoker(invoker IServerCSCmds, defaultInvoker func(*rpc.SimpleCmd, *tcp.Session)) *ServerCSCmdsInvoker {
	inv := new(ServerCSCmdsInvoker)
	inv.invoker = invoker
	inv.defaultInvoker = defaultInvoker
	return inv
} 

func (this *ServerCSCmdsInvoker) Invoke(cmd *rpc.SimpleCmd, se *tcp.Session) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	pack := cmd.Pack()
	switch(cmd.Opcode()) {
	case 22001: 
		this.invoker.GW2CS_Ping(cmd,se)
	case 22004: 
		this.invoker.GW2CS_LoginGuest(cmd,se, pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString())
	
	default:
		if this.defaultInvoker != nil {
			this.defaultInvoker(cmd,se)
		}
	}
	return nil
}

