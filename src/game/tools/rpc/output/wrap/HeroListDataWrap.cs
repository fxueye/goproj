using System;
using Common.Net.Simple;

namespace Game 
{
	public class HeroListDataWrap : IWrapper
	{	
		public HeroDataWrap[] HerosList; // 英雄列表
		
		public void Decode(Packet pck)
		{	
			HerosList = new HeroDataWrap[pck.GetShort()];
			for (int i = 0; i < HerosList.Length; i++)
			{
				HerosList[i] = new HeroDataWrap();
				HerosList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (HerosList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)HerosList.Length);
	        	for(int i = 0; i < HerosList.Length; i++)
	        	{
	        		HerosList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}