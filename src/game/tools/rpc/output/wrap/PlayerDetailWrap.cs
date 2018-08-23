using System;
using Common.Net.Simple;

namespace Game 
{
	public class PlayerDetailWrap : IWrapper
	{	
		public long UID; // 游戏的uid
		public string UserName; // 用户名
		public int Icon; // 头像
		public int GuildID; // 公会ID
		public string GuildName; // 公会名
		public int GuildIcon; // 公会图标
		public short GuildRank; // 公会官阶
		public int Level; // 等级
		public int Exp; // 经验
		public int HeroCnt; // 英雄数量
		public int FavorHero; // 最常用英雄
		public int LadderRank; // 天梯名次
		public int MaxScore; // 天梯最高积分
		public int Score; // 天梯积分
		public int ArenaWin; // 竞技场最高胜场
		public int DonateCnt; // 赠送武将个数
		public BattleHeroWrap[] Heroes; // 上阵英雄
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			UserName = pck.GetString(); 
			Icon = pck.GetInt(); 
			GuildID = pck.GetInt(); 
			GuildName = pck.GetString(); 
			GuildIcon = pck.GetInt(); 
			GuildRank = pck.GetShort(); 
			Level = pck.GetInt(); 
			Exp = pck.GetInt(); 
			HeroCnt = pck.GetInt(); 
			FavorHero = pck.GetInt(); 
			LadderRank = pck.GetInt(); 
			MaxScore = pck.GetInt(); 
			Score = pck.GetInt(); 
			ArenaWin = pck.GetInt(); 
			DonateCnt = pck.GetInt(); 
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
        	pck.PutInt(Icon); 
        	pck.PutInt(GuildID); 
        	pck.PutString(GuildName); 
        	pck.PutInt(GuildIcon); 
        	pck.PutShort(GuildRank); 
        	pck.PutInt(Level); 
        	pck.PutInt(Exp); 
        	pck.PutInt(HeroCnt); 
        	pck.PutInt(FavorHero); 
        	pck.PutInt(LadderRank); 
        	pck.PutInt(MaxScore); 
        	pck.PutInt(Score); 
        	pck.PutInt(ArenaWin); 
        	pck.PutInt(DonateCnt); 
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