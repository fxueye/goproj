using System;
using Common.Net.Simple;

namespace Game 
{
	public class ShopItemListDataWrap : IWrapper
	{	
		public ShopItemDataWrap[] ShopItemList; // 商店item数据
		public long NextTime; // 下次跟新时间
		
		public void Decode(Packet pck)
		{	
			ShopItemList = new ShopItemDataWrap[pck.GetShort()];
			for (int i = 0; i < ShopItemList.Length; i++)
			{
				ShopItemList[i] = new ShopItemDataWrap();
				ShopItemList[i].Decode(pck);
				
			}
			
			NextTime = pck.GetLong(); 
		}
		
        public void Encode(Packet pck)
        {	
        	if (ShopItemList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)ShopItemList.Length);
	        	for(int i = 0; i < ShopItemList.Length; i++)
	        	{
	        		ShopItemList[i].Encode(pck);
					
	        	}
	        }
        	
        	pck.PutLong(NextTime); 
        }
	}
}