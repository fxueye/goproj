using System;
using Common.Net.Simple;

namespace Game 
{
	public class ArenaWrap : IWrapper
	{	
		public BattleHeroWrap[] Heroes; // 可选英雄
		public BattleHeroWrap[] OnBattleHeroes; // 上阵英雄
		
		public void Decode(Packet pck)
		{	
			Heroes = new BattleHeroWrap[pck.GetShort()];
			for (int i = 0; i < Heroes.Length; i++)
			{
				Heroes[i] = new BattleHeroWrap();
				Heroes[i].Decode(pck);
				
			}
			
			OnBattleHeroes = new BattleHeroWrap[pck.GetShort()];
			for (int i = 0; i < OnBattleHeroes.Length; i++)
			{
				OnBattleHeroes[i] = new BattleHeroWrap();
				OnBattleHeroes[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (Heroes == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Heroes.Length);
	        	for(int i = 0; i < Heroes.Length; i++)
	        	{
	        		Heroes[i].Encode(pck);
					
	        	}
	        }
        	
        	if (OnBattleHeroes == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)OnBattleHeroes.Length);
	        	for(int i = 0; i < OnBattleHeroes.Length; i++)
	        	{
	        		OnBattleHeroes[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}