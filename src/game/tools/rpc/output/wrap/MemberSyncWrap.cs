using System;
using Common.Net.Simple;

namespace Game 
{
	public class MemberSyncWrap : IWrapper
	{	
		public long UID; // 玩家UID
		public int GuildID; // 公会ID
		public string GuildName; // 名字
		public short GuildRank; // 公会官职
		public int GuildIcon; // 公会图标
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			GuildID = pck.GetInt(); 
			GuildName = pck.GetString(); 
			GuildRank = pck.GetShort(); 
			GuildIcon = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(UID); 
        	pck.PutInt(GuildID); 
        	pck.PutString(GuildName); 
        	pck.PutShort(GuildRank); 
        	pck.PutInt(GuildIcon); 
        }
	}
}