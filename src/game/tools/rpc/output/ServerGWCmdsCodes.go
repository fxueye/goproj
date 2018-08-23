package cmds
import ()

const (
	ServerGWCmds_HEART_BEAT = 0 // 心跳
	ServerGWCmds_LOGIN_GUEST = 10001 // 登录
	ServerGWCmds_LOGIN_PLATFORM = 10002 // 登录(ptID平台）
	ServerGWCmds_BIND_ACCOUNT = 10003 // 注册用户
	ServerGWCmds_CS_PONG = 20001 // cs_pong
	ServerGWCmds_CS_LOGINRSP = 20002 // cs_login_resp
	ServerGWCmds_BROADCAST = 20003 // 广播消息（不定义参数格式，实际参数与广播的消息协议参数一致）
	ServerGWCmds_KICK_OFFLINE = 20004 // 踢用户下线
	
)