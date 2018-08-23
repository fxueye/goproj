using System;
using Common.Net.Simple;

namespace Game 
{
	public class IntArrayWrap : IWrapper
	{	
		public int[] Value; // 整型数组
		
		public void Decode(Packet pck)
		{	
			Value = new int[pck.GetShort()];
			for (int i = 0; i < Value.Length; i++)
			{
				Value[i] = pck.GetInt();
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (Value == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Value.Length);
	        	for(int i = 0; i < Value.Length; i++)
	        	{
	        		pck.PutInt(Value[i]);
	        	}
	        }
        	
        }
	}
}