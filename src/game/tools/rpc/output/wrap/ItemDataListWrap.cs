using System;
using Common.Net.Simple;

namespace Game 
{
	public class ItemDataListWrap : IWrapper
	{	
		public ItemDataWrap[] ItemList; // 道具列表
		
		public void Decode(Packet pck)
		{	
			ItemList = new ItemDataWrap[pck.GetShort()];
			for (int i = 0; i < ItemList.Length; i++)
			{
				ItemList[i] = new ItemDataWrap();
				ItemList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (ItemList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)ItemList.Length);
	        	for(int i = 0; i < ItemList.Length; i++)
	        	{
	        		ItemList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}