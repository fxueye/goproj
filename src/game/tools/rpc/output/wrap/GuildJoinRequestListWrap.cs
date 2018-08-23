using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildJoinRequestListWrap : IWrapper
	{	
		public GuildJoinRequestWrap[] Requests; // 申请列表
		
		public void Decode(Packet pck)
		{	
			Requests = new GuildJoinRequestWrap[pck.GetShort()];
			for (int i = 0; i < Requests.Length; i++)
			{
				Requests[i] = new GuildJoinRequestWrap();
				Requests[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (Requests == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Requests.Length);
	        	for(int i = 0; i < Requests.Length; i++)
	        	{
	        		Requests[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}