using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildMemberListWrap : IWrapper
	{	
		public int GuildID; // 公会id
		public GuildMemberBriefWrap[] Members; // 成员
		
		public void Decode(Packet pck)
		{	
			GuildID = pck.GetInt(); 
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