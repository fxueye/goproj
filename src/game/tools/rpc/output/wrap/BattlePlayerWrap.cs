using System;
using Common.Net.Simple;

namespace Game 
{
	public class BattlePlayerWrap : IWrapper
	{	
		public long UID; // 游戏的uid
		public string UserName; // 用户名（唯一，初始和UID一致）
		public int Score; // 积分
		public int Level; // 等级
		public int Icon; // 头像
		public int StartHP; // 初始血量
		public BattleHeroWrap[] Heroes; // 上阵英雄
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			UserName = pck.GetString(); 
			Score = pck.GetInt(); 
			Level = pck.GetInt(); 
			Icon = pck.GetInt(); 
			StartHP = pck.GetInt(); 
			Heroes = new BattleHeroWrap[pck.GetShort()];
			for (int i = 0; i < Heroes.Length; i++)
			{
				Heroes[i] = new BattleHeroWrap();
				Heroes[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(UID); 
        	pck.PutString(UserName); 
        	pck.PutInt(Score); 
        	pck.PutInt(Level); 
        	pck.PutInt(Icon); 
        	pck.PutInt(StartHP); 
        	if (Heroes == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Heroes.Length);
	        	for(int i = 0; i < Heroes.Length; i++)
	        	{
	        		Heroes[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}