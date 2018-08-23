using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildDemandWrap : IWrapper
	{	
		public long MemberID; // 会员id
		public string Name; // 名字
		public int IconID; // 头像ID
		public long DemandTime; // 请求时间
		public long ExpireTime; // 过期时间
		public int HeroID; // 索要英雄
		public int MaxCount; // 最大数量
		public int CurCount; // 当前数量
		
		public void Decode(Packet pck)
		{	
			MemberID = pck.GetLong(); 
			Name = pck.GetString(); 
			IconID = pck.GetInt(); 
			DemandTime = pck.GetLong(); 
			ExpireTime = pck.GetLong(); 
			HeroID = pck.GetInt(); 
			MaxCount = pck.GetInt(); 
			CurCount = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(MemberID); 
        	pck.PutString(Name); 
        	pck.PutInt(IconID); 
        	pck.PutLong(DemandTime); 
        	pck.PutLong(ExpireTime); 
        	pck.PutInt(HeroID); 
        	pck.PutInt(MaxCount); 
        	pck.PutInt(CurCount); 
        }
	}
}