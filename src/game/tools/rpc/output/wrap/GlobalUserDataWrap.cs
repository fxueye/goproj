using System;
using Common.Net.Simple;

namespace Game 
{
	public class GlobalUserDataWrap : IWrapper
	{	
		public long Uid; // did
		public string Username; // 用户名
		public int Icon; // 头像
		public int Score; // 积分
		public int Level; // 等级
		public long OnlineTime; // 在线时间
		public long SessID; // session
		public bool Reconnect; // 是否重连
		
		public void Decode(Packet pck)
		{	
			Uid = pck.GetLong(); 
			Username = pck.GetString(); 
			Icon = pck.GetInt(); 
			Score = pck.GetInt(); 
			Level = pck.GetInt(); 
			OnlineTime = pck.GetLong(); 
			SessID = pck.GetLong(); 
			Reconnect = pck.GetBool(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(Uid); 
        	pck.PutString(Username); 
        	pck.PutInt(Icon); 
        	pck.PutInt(Score); 
        	pck.PutInt(Level); 
        	pck.PutLong(OnlineTime); 
        	pck.PutLong(SessID); 
        	pck.PutBool(Reconnect); 
        }
	}
}