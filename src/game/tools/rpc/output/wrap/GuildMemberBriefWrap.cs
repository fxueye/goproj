using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildMemberBriefWrap : IWrapper
	{	
		public long UID; // 玩家UID
		public string Name; // 名字
		public int Level; // 等级
		public int Score; // 积分
		public short GuildRank; // 公会官职
		public int Donation; // 捐献数
		public bool Online; // 是否在线
		public int Icon; // 头像
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			Name = pck.GetString(); 
			Level = pck.GetInt(); 
			Score = pck.GetInt(); 
			GuildRank = pck.GetShort(); 
			Donation = pck.GetInt(); 
			Online = pck.GetBool(); 
			Icon = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(UID); 
        	pck.PutString(Name); 
        	pck.PutInt(Level); 
        	pck.PutInt(Score); 
        	pck.PutShort(GuildRank); 
        	pck.PutInt(Donation); 
        	pck.PutBool(Online); 
        	pck.PutInt(Icon); 
        }
	}
}