using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildBriefWrap : IWrapper
	{	
		public int GuildID; // 公会id
		public string Name; // 名字
		public int IconID; // 头像ID
		public short Member; // 成员数
		public int Score; // 公会积分
		
		public void Decode(Packet pck)
		{	
			GuildID = pck.GetInt(); 
			Name = pck.GetString(); 
			IconID = pck.GetInt(); 
			Member = pck.GetShort(); 
			Score = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(GuildID); 
        	pck.PutString(Name); 
        	pck.PutInt(IconID); 
        	pck.PutShort(Member); 
        	pck.PutInt(Score); 
        }
	}
}