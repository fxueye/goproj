using System;
using Common.Net.Simple;

namespace Game 
{
	public class ChestDataWrap : IWrapper
	{	
		public int ChestIdx; // 宝箱位置
		public int ChestID; // 宝箱ID
		public int QuickenCount; // 加速次数
		public long ActiveTime; // 激活时间
		public int Lv; // 宝箱等级
		
		public void Decode(Packet pck)
		{	
			ChestIdx = pck.GetInt(); 
			ChestID = pck.GetInt(); 
			QuickenCount = pck.GetInt(); 
			ActiveTime = pck.GetLong(); 
			Lv = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ChestIdx); 
        	pck.PutInt(ChestID); 
        	pck.PutInt(QuickenCount); 
        	pck.PutLong(ActiveTime); 
        	pck.PutInt(Lv); 
        }
	}
}