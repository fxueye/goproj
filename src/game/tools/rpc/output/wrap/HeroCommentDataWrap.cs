using System;
using Common.Net.Simple;

namespace Game 
{
	public class HeroCommentDataWrap : IWrapper
	{	
		public int CommentID; // 评论ID
		public int HeroID; // 英雄ID
		public long UserID; // 用户ID
		public string UserNick; // 用户昵称
		public string Content; // 内容
		public int LikeNum; // 点赞数
		public long CommentTime; // 评论时间
		public bool IsLiked; // 是否已经点赞
		
		public void Decode(Packet pck)
		{	
			CommentID = pck.GetInt(); 
			HeroID = pck.GetInt(); 
			UserID = pck.GetLong(); 
			UserNick = pck.GetString(); 
			Content = pck.GetString(); 
			LikeNum = pck.GetInt(); 
			CommentTime = pck.GetLong(); 
			IsLiked = pck.GetBool(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(CommentID); 
        	pck.PutInt(HeroID); 
        	pck.PutLong(UserID); 
        	pck.PutString(UserNick); 
        	pck.PutString(Content); 
        	pck.PutInt(LikeNum); 
        	pck.PutLong(CommentTime); 
        	pck.PutBool(IsLiked); 
        }
	}
}