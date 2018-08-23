using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildJoinRequestWrap : IWrapper
	{	
		public long UID; // 玩家UID
		public string Name; // 名字
		public int Level; // 等级
		public int Score; // 积分
		public int Icon; // 头像
		public string Msg; // 留言
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			Name = pck.GetString(); 
			Level = pck.GetInt(); 
			Score = pck.GetInt(); 
			Icon = pck.GetInt(); 
			Msg = pck.GetString(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(UID); 
        	pck.PutString(Name); 
        	pck.PutInt(Level); 
        	pck.PutInt(Score); 
        	pck.PutInt(Icon); 
        	pck.PutString(Msg); 
        }
	}
}