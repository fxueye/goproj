using System;
using Common.Net.Simple;

namespace Game 
{
	public class PlayerDataWrap : IWrapper
	{	
		public string[] DataKey; // 数据Key
		public long[] DataValue; // 数据Value
		
		public void Decode(Packet pck)
		{	
			DataKey = new string[pck.GetShort()];
			for (int i = 0; i < DataKey.Length; i++)
			{
				DataKey[i] = pck.GetString();
			}
			
			DataValue = new long[pck.GetShort()];
			for (int i = 0; i < DataValue.Length; i++)
			{
				DataValue[i] = pck.GetLong();
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (DataKey == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)DataKey.Length);
	        	for(int i = 0; i < DataKey.Length; i++)
	        	{
	        		pck.PutString(DataKey[i]);
	        	}
	        }
        	
        	if (DataValue == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)DataValue.Length);
	        	for(int i = 0; i < DataValue.Length; i++)
	        	{
	        		pck.PutLong(DataValue[i]);
	        	}
	        }
        	
        }
	}
}