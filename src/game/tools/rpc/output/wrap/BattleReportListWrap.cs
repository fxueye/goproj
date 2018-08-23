using System;
using Common.Net.Simple;

namespace Game 
{
	public class BattleReportListWrap : IWrapper
	{	
		public BattleReportWrap[] List; // 列表
		
		public void Decode(Packet pck)
		{	
			List = new BattleReportWrap[pck.GetShort()];
			for (int i = 0; i < List.Length; i++)
			{
				List[i] = new BattleReportWrap();
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