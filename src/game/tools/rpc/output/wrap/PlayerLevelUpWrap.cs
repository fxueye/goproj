using System;
using Common.Net.Simple;

namespace Game 
{
	public class PlayerLevelUpWrap : IWrapper
	{	
		public int oldLv; // 旧等级
		public int newLv; // 新等级
		public int oldBlood; // 旧血量
		public int newBlood; // 新血量
		public int diamondAdd; // 获得钻石
		public int coinAdd; // 获得金币
		
		public void Decode(Packet pck)
		{	
			oldLv = pck.GetInt(); 
			newLv = pck.GetInt(); 
			oldBlood = pck.GetInt(); 
			newBlood = pck.GetInt(); 
			diamondAdd = pck.GetInt(); 
			coinAdd = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(oldLv); 
        	pck.PutInt(newLv); 
        	pck.PutInt(oldBlood); 
        	pck.PutInt(newBlood); 
        	pck.PutInt(diamondAdd); 
        	pck.PutInt(coinAdd); 
        }
	}
}