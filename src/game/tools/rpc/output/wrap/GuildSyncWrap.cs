using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildSyncWrap : IWrapper
	{	
		public int GuildID; // 公会ID
		public string GuildName; // 名字
		public int GuildIcon; // 公会图标
		
		public void Decode(Packet pck)
		{	
			GuildID = pck.GetInt(); 
			GuildName = pck.GetString(); 
			GuildIcon = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(GuildID); 
        	pck.PutString(GuildName); 
        	pck.PutInt(GuildIcon); 
        }
	}
}