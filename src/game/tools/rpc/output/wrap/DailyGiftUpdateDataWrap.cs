using System;
using Common.Net.Simple;

namespace Game 
{
	public class DailyGiftUpdateDataWrap : IWrapper
	{	
		public int GiftIdx; // 当前领取的idx
		public int Status; // 更新状态
		
		public void Decode(Packet pck)
		{	
			GiftIdx = pck.GetInt(); 
			Status = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(GiftIdx); 
        	pck.PutInt(Status); 
        }
	}
}