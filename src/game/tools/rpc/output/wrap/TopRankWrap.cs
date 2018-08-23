using System;
using Common.Net.Simple;

namespace Game 
{
	public class TopRankWrap : IWrapper
	{	
		public long UID; // UID
		public string Name; // 名字
		public int IconID; // 头像ID
		public int[] Data; // 数据
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			Name = pck.GetString(); 
			IconID = pck.GetInt(); 
			Data = new int[pck.GetShort()];
			for (int i = 0; i < Data.Length; i++)
			{
				Data[i] = pck.GetInt();
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(UID); 
        	pck.PutString(Name); 
        	pck.PutInt(IconID); 
        	if (Data == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Data.Length);
	        	for(int i = 0; i < Data.Length; i++)
	        	{
	        		pck.PutInt(Data[i]);
	        	}
	        }
        	
        }
	}
}