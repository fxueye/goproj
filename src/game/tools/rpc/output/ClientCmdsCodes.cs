namespace Game 
{
	public class ClientCmdsCodes
	{
		public const short HEART_BEAT = (short)0; // 心跳
		public const short LOGIN_SUCCESS = (short)1; // 登录成功
		public const short LOGIN_FAILED = (short)2; // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
		
	}
}