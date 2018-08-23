using System;
using Common.Net.Simple;

namespace Game 
{
	public class QuestDataWrap : IWrapper
	{	
		public int QuestID; // 关卡ID
		public bool ChestState; // 宝箱是否发放
		public long LastUpdateTime; // 最后一次更新时间
		
		public void Decode(Packet pck)
		{	
			QuestID = pck.GetInt(); 
			ChestState = pck.GetBool(); 
			LastUpdateTime = pck.GetLong(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(QuestID); 
        	pck.PutBool(ChestState); 
        	pck.PutLong(LastUpdateTime); 
        }
	}
}