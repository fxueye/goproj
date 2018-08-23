using System;
using Common.Net.Simple;

namespace Game 
{
	public class BattleHeroWrap : IWrapper
	{	
		public int HeroID; // 英雄ID
		public int Level; // 英雄等级
		
		public void Decode(Packet pck)
		{	
			HeroID = pck.GetInt(); 
			Level = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(HeroID); 
        	pck.PutInt(Level); 
        }
	}
}