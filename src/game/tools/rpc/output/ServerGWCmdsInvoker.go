package cmds
import (
	"errors"
	"fmt"
	rpc "tipcat.com/common/rpc/simple"
	tcp "tipcat.com/common/server/tcp"
	
)

type IServerGWCmds interface {
	HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session) // 心跳
	LoginGuest(cmd *rpc.SimpleCmd, se *tcp.Session, devID string, deviceType string, partnerID string, version string) // 登录
	LoginPlatform(cmd *rpc.SimpleCmd, se *tcp.Session, ptID string, account string, deviceType string, partnerID string, version string, reconnect bool, token string, extension string) // 登录(ptID平台）
	BindAccount(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, ptID string, ptData string) // 注册用户
	CS_Pong(cmd *rpc.SimpleCmd, se *tcp.Session) // cs_pong
	CS_LoginRsp(cmd *rpc.SimpleCmd, se *tcp.Session, suc bool, uid int64, accName string) // cs_login_resp
	Broadcast(cmd *rpc.SimpleCmd, se *tcp.Session) // 广播消息（不定义参数格式，实际参数与广播的消息协议参数一致）
	KickOffLine(cmd *rpc.SimpleCmd, se *tcp.Session, sid int64) // 踢用户下线
	
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
	case 10003: 
		this.invoker.BindAccount(cmd,se, pack.PopInt64(), pack.PopString(), pack.PopString())
	case 20001: 
		this.invoker.CS_Pong(cmd,se)
	case 20002: 
		this.invoker.CS_LoginRsp(cmd,se, pack.PopBool(), pack.PopInt64(), pack.PopString())
	case 20003: 
		this.invoker.Broadcast(cmd,se)
	case 20004: 
		this.invoker.KickOffLine(cmd,se, pack.PopInt64())
	
	default:
		if this.defaultInvoker != nil {
			this.defaultInvoker(cmd,se)
		}
	}
	return nil
}

