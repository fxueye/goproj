using System;
using Common.Net.Simple;

namespace Game 
{
	public class ActivityDungeonWrap : IWrapper
	{	
		public int GroupID; // 活动本groupid
		public long StartTime; // 开始时间
		public long EndTime; // 结束时间
		
		public void Decode(Packet pck)
		{	
			GroupID = pck.GetInt(); 
			StartTime = pck.GetLong(); 
			EndTime = pck.GetLong(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(GroupID); 
        	pck.PutLong(StartTime); 
        	pck.PutLong(EndTime); 
        }
	}
}