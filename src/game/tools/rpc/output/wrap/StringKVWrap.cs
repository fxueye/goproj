using System;
using Common.Net.Simple;

namespace Game 
{
	public class StringKVWrap : IWrapper
	{	
		public string Key; // map的key
		public string Value; // map的value
		
		public void Decode(Packet pck)
		{	
			Key = pck.GetString(); 
			Value = pck.GetString(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutString(Key); 
        	pck.PutString(Value); 
        }
	}
}