using System;
using Common.Net.Simple;

namespace Game 
{
	public class IntKVWrap : IWrapper
	{	
		public short Key; // map的key
		public short Value; // map的value
		
		public void Decode(Packet pck)
		{	
			Key = pck.GetShort(); 
			Value = pck.GetShort(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutShort(Key); 
        	pck.PutShort(Value); 
        }
	}
}