using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildDetailWrap : IWrapper
	{	
		public int GuildID; // 公会id
		public string GuildName; // 名字
		public int IconID; // 头像ID
		public int GuildScore; // 公会积分
		public string GuildNote; // 公会公告
		public short JoinType; // 加入类型
		public int JoinScore; // 加入积分
		public int TopRank; // 公会排名
		public int WeekDonate; // 公会本周捐献
		public GuildMemberBriefWrap[] Members; // 成员
		
		public void Decode(Packet pck)
		{	
			GuildID = pck.GetInt(); 
			GuildName = pck.GetString(); 
			IconID = pck.GetInt(); 
			GuildScore = pck.GetInt(); 
			GuildNote = pck.GetString(); 
			JoinType = pck.GetShort(); 
			JoinScore = pck.GetInt(); 
			TopRank = pck.GetInt(); 
			WeekDonate = pck.GetInt(); 
			Members = new GuildMemberBriefWrap[pck.GetShort()];
			for (int i = 0; i < Members.Length; i++)
			{
				Members[i] = new GuildMemberBriefWrap();
				Members[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(GuildID); 
        	pck.PutString(GuildName); 
        	pck.PutInt(IconID); 
        	pck.PutInt(GuildScore); 
        	pck.PutString(GuildNote); 
        	pck.PutShort(JoinType); 
        	pck.PutInt(JoinScore); 
        	pck.PutInt(TopRank); 
        	pck.PutInt(WeekDonate); 
        	if (Members == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Members.Length);
	        	for(int i = 0; i < Members.Length; i++)
	        	{
	        		Members[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}