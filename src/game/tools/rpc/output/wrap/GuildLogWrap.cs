using System;
using Common.Net.Simple;

namespace Game 
{
	public class GuildLogWrap : IWrapper
	{	
		public short Opcode; // 操作类型
		public long CreateTime; // 时间
		public long MemberID; // 操作者ID
		public string MemberName; // 操作者名字
		public long TargetID; // 目标ID
		public string TargetName; // 目标名字
		public short GuildRank; // 操作官职
		
		public void Decode(Packet pck)
		{	
			Opcode = pck.GetShort(); 
			CreateTime = pck.GetLong(); 
			MemberID = pck.GetLong(); 
			MemberName = pck.GetString(); 
			TargetID = pck.GetLong(); 
			TargetName = pck.GetString(); 
			GuildRank = pck.GetShort(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutShort(Opcode); 
        	pck.PutLong(CreateTime); 
        	pck.PutLong(MemberID); 
        	pck.PutString(MemberName); 
        	pck.PutLong(TargetID); 
        	pck.PutString(TargetName); 
        	pck.PutShort(GuildRank); 
        }
	}
}