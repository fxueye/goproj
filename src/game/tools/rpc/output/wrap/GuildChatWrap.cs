using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildChatWrap : IWrapper
	{	
		public long MemberID; // 会员id
		public string Name; // 名字
		public int IconID; // 头像ID
		public long ChatTime; // 聊天时间
		public string Msg; // 聊天内容
		
		public void Decode(Packet pck)
		{	
			MemberID = pck.GetLong(); 
			Name = pck.GetString(); 
			IconID = pck.GetInt(); 
			ChatTime = pck.GetLong(); 
			Msg = pck.GetString(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(MemberID); 
        	pck.PutString(Name); 
        	pck.PutInt(IconID); 
        	pck.PutLong(ChatTime); 
        	pck.PutString(Msg); 
        }
	}
}