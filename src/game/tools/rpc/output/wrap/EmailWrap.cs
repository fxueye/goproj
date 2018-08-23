using System;
using Common.Net.Simple;

namespace Game 
{
	public class EmailWrap : IWrapper
	{	
		public int ID; // 邮件id
		public string Title; // 邮件标题
		public string Content; // 邮件内容
		public short Flag; // 邮件是否读取
		public string Attachment; // 邮件附件
		public short AttachmentState; // 邮件附件是否已经领取
		public long EmailTime; // 邮件创建时间
		public string From; // 邮件发送者
		
		public void Decode(Packet pck)
		{	
			ID = pck.GetInt(); 
			Title = pck.GetString(); 
			Content = pck.GetString(); 
			Flag = pck.GetShort(); 
			Attachment = pck.GetString(); 
			AttachmentState = pck.GetShort(); 
			EmailTime = pck.GetLong(); 
			From = pck.GetString(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ID); 
        	pck.PutString(Title); 
        	pck.PutString(Content); 
        	pck.PutShort(Flag); 
        	pck.PutString(Attachment); 
        	pck.PutShort(AttachmentState); 
        	pck.PutLong(EmailTime); 
        	pck.PutString(From); 
        }
	}
}