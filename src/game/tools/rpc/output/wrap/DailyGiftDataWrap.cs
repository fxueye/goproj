using System;
using Common.Net.Simple;

namespace Game 
{
	public class DailyGiftDataWrap : IWrapper
	{	
		public int GiftIdx; // 数据下标
		public int GiftID; // 礼品ID
		public int GiftType; // 类型 1是item 2是hero
		public int Count; // 数量
		public int State; // 状态 0已经领取 1可以领取 2不可以领取 3 补签  4 已经补签
		public int GroupID; // 组ID
		
		public void Decode(Packet pck)
		{	
			GiftIdx = pck.GetInt(); 
			GiftID = pck.GetInt(); 
			GiftType = pck.GetInt(); 
			Count = pck.GetInt(); 
			State = pck.GetInt(); 
			GroupID = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(GiftIdx); 
        	pck.PutInt(GiftID); 
        	pck.PutInt(GiftType); 
        	pck.PutInt(Count); 
        	pck.PutInt(State); 
        	pck.PutInt(GroupID); 
        }
	}
}