using System;
using Common.Net.Simple;

namespace Game 
{
	public class AchievementWrap : IWrapper
	{	
		public int AchType; // 成就类型
		public int AchId; // 成就id
		public int AchScore; // 成就得分
		public int AchState; // 成就状态  0进行中 1完成 2领取
		public long LastTime; // 成就最后更新时间
		
		public void Decode(Packet pck)
		{	
			AchType = pck.GetInt(); 
			AchId = pck.GetInt(); 
			AchScore = pck.GetInt(); 
			AchState = pck.GetInt(); 
			LastTime = pck.GetLong(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(AchType); 
        	pck.PutInt(AchId); 
        	pck.PutInt(AchScore); 
        	pck.PutInt(AchState); 
        	pck.PutLong(LastTime); 
        }
	}
}