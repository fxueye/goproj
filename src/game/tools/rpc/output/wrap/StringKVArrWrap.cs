using System;
using Common.Net.Simple;

namespace Game 
{
	public class StringKVArrWrap : IWrapper
	{	
		public StringKVWrap[] StringMap; // map<string,string>
		
		public void Decode(Packet pck)
		{	
			StringMap = new StringKVWrap[pck.GetShort()];
			for (int i = 0; i < StringMap.Length; i++)
			{
				StringMap[i] = new StringKVWrap();
				StringMap[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (StringMap == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)StringMap.Length);
	        	for(int i = 0; i < StringMap.Length; i++)
	        	{
	        		StringMap[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}