using System;
using Common.Net.Simple;

namespace Game 
{
	public class ChestOpenDataWrap : IWrapper
	{	
		public int ChestIdx; // 宝箱位置
		public int ChestID; // 宝箱ID
		public int Coin; // 金币数量
		public int Diamond; // 钻石数量
		public HeroDataWrap[] Heros; // 英雄
		public int LadderLv; // 天体等级
		
		public void Decode(Packet pck)
		{	
			ChestIdx = pck.GetInt(); 
			ChestID = pck.GetInt(); 
			Coin = pck.GetInt(); 
			Diamond = pck.GetInt(); 
			Heros = new HeroDataWrap[pck.GetShort()];
			for (int i = 0; i < Heros.Length; i++)
			{
				Heros[i] = new HeroDataWrap();
				Heros[i].Decode(pck);
				
			}
			
			LadderLv = pck.GetInt(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ChestIdx); 
        	pck.PutInt(ChestID); 
        	pck.PutInt(Coin); 
        	pck.PutInt(Diamond); 
        	if (Heros == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Heros.Length);
	        	for(int i = 0; i < Heros.Length; i++)
	        	{
	        		Heros[i].Encode(pck);
					
	        	}
	        }
        	
        	pck.PutInt(LadderLv); 
        }
	}
}