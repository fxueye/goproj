using System;
using Common.Net.Simple;

namespace Game 
{
	public class HeroCommentListDataWrap : IWrapper
	{	
		public HeroCommentDataWrap[] CommentsList; // 评论列表
		
		public void Decode(Packet pck)
		{	
			CommentsList = new HeroCommentDataWrap[pck.GetShort()];
			for (int i = 0; i < CommentsList.Length; i++)
			{
				CommentsList[i] = new HeroCommentDataWrap();
				CommentsList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (CommentsList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)CommentsList.Length);
	        	for(int i = 0; i < CommentsList.Length; i++)
	        	{
	        		CommentsList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}