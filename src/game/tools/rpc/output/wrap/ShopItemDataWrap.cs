using System;
using Common.Net.Simple;

namespace Game 
{
	public class ShopItemDataWrap : IWrapper
	{	
		public int ShopItemID; // 商品道具ID
		public int GoodID; // 商品id 钻石 金币 宝箱 随机英雄 固定英雄
		public int GoodType; // 商品类型 1 买钻石 2 买金币 3 买宝箱 4 买随机英雄 5 买固定英雄
		public int GoodCount; // 商品数量
		public long BuyTime; // 购买时间
		public int Count; // 商品购买数量
		
		public void Decode(Packet pck)
		{	
			ShopItemID = pck.GetInt(); 
			GoodID = pck.GetInt(); 
			GoodType = pck.GetInt(); 
			GoodCount = pck.GetInt(); 
			BuyTime = pck.GetLong(); 
			Count = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ShopItemID); 
        	pck.PutInt(GoodID); 
        	pck.PutInt(GoodType); 
        	pck.PutInt(GoodCount); 
        	pck.PutLong(BuyTime); 
        	pck.PutInt(Count); 
        }
	}
}