using System;
using Common.Net.Simple;

namespace Game 
{
	public class DailyGiftListDataWrap : IWrapper
	{	
		public DailyGiftDataWrap[] DailyList; // 日常数据
		public long NextTime; // 下次领取新时间
		public DailyActivityDataWrap[] ActivityList; // 活跃度数据
		
		public void Decode(Packet pck)
		{	
			DailyList = new DailyGiftDataWrap[pck.GetShort()];
			for (int i = 0; i < DailyList.Length; i++)
			{
				DailyList[i] = new DailyGiftDataWrap();
				DailyList[i].Decode(pck);
				
			}
			
			NextTime = pck.GetLong(); 
			ActivityList = new DailyActivityDataWrap[pck.GetShort()];
			for (int i = 0; i < ActivityList.Length; i++)
			{
				ActivityList[i] = new DailyActivityDataWrap();
				ActivityList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (DailyList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)DailyList.Length);
	        	for(int i = 0; i < DailyList.Length; i++)
	        	{
	        		DailyList[i].Encode(pck);
					
	        	}
	        }
        	
        	pck.PutLong(NextTime); 
        	if (ActivityList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)ActivityList.Length);
	        	for(int i = 0; i < ActivityList.Length; i++)
	        	{
	        		ActivityList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}