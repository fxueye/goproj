using System;
using Common.Net.Simple;

namespace Game 
{
	public class DailyGetActivityDataWrap : IWrapper
	{	
		public int[] ActIdxs; // 当前领取的id数组
		public ChestOpenDataWrap ChestDataWrap; // 领取的结果数据
		
		public void Decode(Packet pck)
		{	
			ActIdxs = new int[pck.GetShort()];
			for (int i = 0; i < ActIdxs.Length; i++)
			{
				ActIdxs[i] = pck.GetInt();
			}
			
			ChestDataWrap = new ChestOpenDataWrap();
			ChestDataWrap.Decode(pck);
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (ActIdxs == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)ActIdxs.Length);
	        	for(int i = 0; i < ActIdxs.Length; i++)
	        	{
	        		pck.PutInt(ActIdxs[i]);
	        	}
	        }
        	
        	ChestDataWrap.Encode(pck); 
        }
	}
}