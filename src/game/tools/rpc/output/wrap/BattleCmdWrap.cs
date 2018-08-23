using System;
using Common.Net.Simple;

namespace Game 
{
	public class BattleCmdWrap : IWrapper
	{	
		public int SeqID; // 序列号
		public short Opcode; // 协议号
		public long Timestamp; // 时间戳
		public int[] Args; // 参数数组
		
		public void Decode(Packet pck)
		{	
			SeqID = pck.GetInt(); 
			Opcode = pck.GetShort(); 
			Timestamp = pck.GetLong(); 
			Args = new int[pck.GetShort()];
			for (int i = 0; i < Args.Length; i++)
			{
				Args[i] = pck.GetInt();
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(SeqID); 
        	pck.PutShort(Opcode); 
        	pck.PutLong(Timestamp); 
        	if (Args == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Args.Length);
	        	for(int i = 0; i < Args.Length; i++)
	        	{
	        		pck.PutInt(Args[i]);
	        	}
	        }
        	
        }
	}
}