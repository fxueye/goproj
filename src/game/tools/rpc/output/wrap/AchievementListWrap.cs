using System;
using Common.Net.Simple;

namespace Game 
{
	public class AchievementListWrap : IWrapper
	{	
		public AchievementWrap[] AchList; // 成就列表
		
		public void Decode(Packet pck)
		{	
			AchList = new AchievementWrap[pck.GetShort()];
			for (int i = 0; i < AchList.Length; i++)
			{
				AchList[i] = new AchievementWrap();
				AchList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (AchList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)AchList.Length);
	        	for(int i = 0; i < AchList.Length; i++)
	        	{
	        		AchList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}