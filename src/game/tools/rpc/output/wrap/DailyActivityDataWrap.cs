using System;
using Common.Net.Simple;

namespace Game 
{
	public class DailyActivityDataWrap : IWrapper
	{	
		public int ActivityIdx; // 数据下标
		public int ChestID; // 宝箱ID
		public int ActivityPoints; // 活跃度
		public int Lv; // 等级
		public int State; // 状态 0已经领取 1可以领取 2不可以领取
		
		public void Decode(Packet pck)
		{	
			ActivityIdx = pck.GetInt(); 
			ChestID = pck.GetInt(); 
			ActivityPoints = pck.GetInt(); 
			Lv = pck.GetInt(); 
			State = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ActivityIdx); 
        	pck.PutInt(ChestID); 
        	pck.PutInt(ActivityPoints); 
        	pck.PutInt(Lv); 
        	pck.PutInt(State); 
        }
	}
}