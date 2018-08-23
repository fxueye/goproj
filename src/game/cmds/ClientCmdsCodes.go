package cmds
import ()

const (
	ClientCmds_HEART_BEAT = 0 // 心跳
	ClientCmds_LOGIN_SUCCESS = 1 // 登录成功
	ClientCmds_LOGIN_FAILED = 2 // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
	
)