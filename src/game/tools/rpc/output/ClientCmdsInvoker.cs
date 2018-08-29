using System;
using Common.Log;
using Common.Net.Simple;

namespace Game
{
	public interface IClientCmds 
	{
		void HeartBeat(Command cmd, string msg); // 心跳
		void LoginSuccess(Command cmd, PlayerWrap player, bool reconnect, string extension); // 登录成功
		void LoginFailed(Command cmd, short errorCode, string errMsg); // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
		
	}
	
    public class ClientCmdsInvoker : IInvoker
    {
        IClientCmds _cmds = null;
		Action<short> _onCmdInvoked = null;
        public ClientCmdsInvoker(IClientCmds cmds)
        {
            _cmds = cmds;
        }
		public void SetOnCmdInvoked(Action<short> onCmdInvoked)
		{
			_onCmdInvoked = onCmdInvoked;
		}

        public void Invoke(Command cmd)
        {
        	try
        	{
        		Packet pack = cmd.Pack;
	        	switch (cmd.Opcode)
	            {
	            	case (short)0: _cmds.HeartBeat(cmd, pack.GetString()); break;
	            	case (short)1: _cmds.LoginSuccess(cmd, new PlayerWrap().Decode(), pack.GetBool(), pack.GetString()); break;
	            	case (short)2: _cmds.LoginFailed(cmd, pack.GetShort(), pack.GetString()); break;
	            	
	            }
				if (_onCmdInvoked != null)
					_onCmdInvoked(cmd.Opcode);
        	}
        	catch(Exception e)
        	{
				L.Error("invoke error, opcode=" + cmd.Opcode);
        		L.Exception(e.Message, e);
        	}
		}
    }
}
