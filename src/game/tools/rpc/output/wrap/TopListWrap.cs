using System;
using Common.Net.Simple;

namespace Game 
{
	public class TopListWrap : IWrapper
	{	
		public short Kind; // 类型
		public TopRankWrap[] List; // 数据
		
		public void Decode(Packet pck)
		{	
			Kind = pck.GetShort(); 
			List = new TopRankWrap[pck.GetShort()];
			for (int i = 0; i < List.Length; i++)
			{
				List[i] = new TopRankWrap();
				List[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutShort(Kind); 
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