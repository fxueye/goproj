using System;
using Common.Net.Simple;

namespace Game 
{
	public class ChestListDataWrap : IWrapper
	{	
		public ChestDataWrap[] ChestList; // 宝箱列表
		
		public void Decode(Packet pck)
		{	
			ChestList = new ChestDataWrap[pck.GetShort()];
			for (int i = 0; i < ChestList.Length; i++)
			{
				ChestList[i] = new ChestDataWrap();
				ChestList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (ChestList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)ChestList.Length);
	        	for(int i = 0; i < ChestList.Length; i++)
	        	{
	        		ChestList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}