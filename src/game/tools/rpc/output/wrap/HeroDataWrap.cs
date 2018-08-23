using System;
using Common.Net.Simple;

namespace Game 
{
	public class HeroDataWrap : IWrapper
	{	
		public int HeroID; // 卡牌ID
		public int Count; // 卡牌数量
		public int Level; // 等级
		public long LastUpdateTime; // 最后一次更新时间
		public int UseCnt; // 使用次数
		
		public void Decode(Packet pck)
		{	
			HeroID = pck.GetInt(); 
			Count = pck.GetInt(); 
			Level = pck.GetInt(); 
			LastUpdateTime = pck.GetLong(); 
			UseCnt = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(HeroID); 
        	pck.PutInt(Count); 
        	pck.PutInt(Level); 
        	pck.PutLong(LastUpdateTime); 
        	pck.PutInt(UseCnt); 
        }
	}
}