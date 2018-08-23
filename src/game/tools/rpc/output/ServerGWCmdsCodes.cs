namespace Game 
{
	public class ServerGWCmdsCodes
	{
		public const short HEART_BEAT = (short)0; // 心跳
		public const short LOGIN_GUEST = (short)10001; // 登录
		public const short LOGIN_PLATFORM = (short)10002; // 登录(ptID平台）
		public const short BIND_ACCOUNT = (short)10003; // 注册用户
		public const short CS_PONG = (short)20001; // cs_pong
		public const short CS_LOGINRSP = (short)20002; // cs_login_resp
		public const short BROADCAST = (short)20003; // 广播消息（不定义参数格式，实际参数与广播的消息协议参数一致）
		public const short KICK_OFFLINE = (short)20004; // 踢用户下线
		
	}
}