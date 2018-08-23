using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildPvPRequestWrap : IWrapper
	{	
		public long MemberID; // 会员id
		public string Name; // 名字
		public int IconID; // 头像ID
		public string Msg; // 留言内容
		
		public void Decode(Packet pck)
		{	
			MemberID = pck.GetLong(); 
			Name = pck.GetString(); 
			IconID = pck.GetInt(); 
			Msg = pck.GetString(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(MemberID); 
        	pck.PutString(Name); 
        	pck.PutInt(IconID); 
        	pck.PutString(Msg); 
        }
	}
}