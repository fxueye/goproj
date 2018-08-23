using System;
using Common.Net.Simple;

namespace Game 
{
	public class ProductInfoWrap : IWrapper
	{	
		public string ProductID; // 商品ID
		public short SortID; // 排序ID
		public string ProductName; // 商品名
		public string Desc; // 描述
		public bool Visible; // 是否可见
		public bool Hot; // 是否热销
		public int MoneyType; // 货币类型
		public int MoneyCnt; // 发货数量
		public int GiftMoneyCnt; // 赠送数量
		public float Price; // 购买价格
		public string Currency; // 货币类型
		
		public void Decode(Packet pck)
		{	
			ProductID = pck.GetString(); 
			SortID = pck.GetShort(); 
			ProductName = pck.GetString(); 
			Desc = pck.GetString(); 
			Visible = pck.GetBool(); 
			Hot = pck.GetBool(); 
			MoneyType = pck.GetInt(); 
			MoneyCnt = pck.GetInt(); 
			GiftMoneyCnt = pck.GetInt(); 
			Price = pck.GetFloat(); 
			Currency = pck.GetString(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutString(ProductID); 
        	pck.PutShort(SortID); 
        	pck.PutString(ProductName); 
        	pck.PutString(Desc); 
        	pck.PutBool(Visible); 
        	pck.PutBool(Hot); 
        	pck.PutInt(MoneyType); 
        	pck.PutInt(MoneyCnt); 
        	pck.PutInt(GiftMoneyCnt); 
        	pck.PutFloat(Price); 
        	pck.PutString(Currency); 
        }
	}
}