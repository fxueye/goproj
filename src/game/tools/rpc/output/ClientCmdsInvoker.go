package cmds
import (
	"errors"
	"fmt"
	rpc "tipcat.com/common/rpc/simple"
	tcp "tipcat.com/common/server/tcp"
	wraps "tipcat.com/cmds/wraps"
)

type IClientCmds interface {
	HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session) // 心跳
	LoginSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, player *wraps.PlayerWrap, reconnect bool, extension string) // 登录成功
	LoginFailed(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int16, errMsg string) // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
	
}

type ClientCmdsInvoker struct {
	invoker IClientCmds
	defaultInvoker func(cmd *rpc.SimpleCmd, se *tcp.Session)
	rpc.SimpleInvoker
} 

func NewClientCmdsInvoker(invoker IClientCmds, defaultInvoker func(*rpc.SimpleCmd, *tcp.Session)) *ClientCmdsInvoker {
	inv := new(ClientCmdsInvoker)
	inv.invoker = invoker
	inv.defaultInvoker = defaultInvoker
	return inv
} 

func (this *ClientCmdsInvoker) Invoke(cmd *rpc.SimpleCmd, se *tcp.Session) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	pack := cmd.Pack()
	switch(cmd.Opcode()) {
	case 0: 
		this.invoker.HeartBeat(cmd,se)
	case 1: 
		this.invoker.LoginSuccess(cmd,se, new(wraps.PlayerWrap).Decode(pack).(*wraps.PlayerWrap), pack.PopBool(), pack.PopString())
	case 2: 
		this.invoker.LoginFailed(cmd,se, pack.PopInt16(), pack.PopString())
	
	default:
		if this.defaultInvoker != nil {
			this.defaultInvoker(cmd,se)
		}
	}
	return nil
}

