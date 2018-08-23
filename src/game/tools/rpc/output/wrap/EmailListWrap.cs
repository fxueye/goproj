using System;
using Common.Net.Simple;

namespace Game 
{
	public class EmailListWrap : IWrapper
	{	
		public EmailWrap[] List; // 数据
		
		public void Decode(Packet pck)
		{	
			List = new EmailWrap[pck.GetShort()];
			for (int i = 0; i < List.Length; i++)
			{
				List[i] = new EmailWrap();
				List[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (List == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)List.Length);
	        	for(int i = 0; i < List.Length; i++)
	        	{
	        		List[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}