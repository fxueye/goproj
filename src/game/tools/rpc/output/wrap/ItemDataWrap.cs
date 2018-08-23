using System;
using Common.Net.Simple;

namespace Game 
{
	public class ItemDataWrap : IWrapper
	{	
		public int ItemID; // 道具ID
		public int Count; // 道具数量
		public long LastUpdateTime; // 最后一次更新时间
		
		public void Decode(Packet pck)
		{	
			ItemID = pck.GetInt(); 
			Count = pck.GetInt(); 
			LastUpdateTime = pck.GetLong(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ItemID); 
        	pck.PutInt(Count); 
        	pck.PutLong(LastUpdateTime); 
        }
	}
}