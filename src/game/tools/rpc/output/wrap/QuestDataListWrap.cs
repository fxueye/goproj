using System;
using Common.Net.Simple;

namespace Game 
{
	public class QuestDataListWrap : IWrapper
	{	
		public QuestDataWrap[] QuestList; // 关卡列表
		
		public void Decode(Packet pck)
		{	
			QuestList = new QuestDataWrap[pck.GetShort()];
			for (int i = 0; i < QuestList.Length; i++)
			{
				QuestList[i] = new QuestDataWrap();
				QuestList[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	if (QuestList == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)QuestList.Length);
	        	for(int i = 0; i < QuestList.Length; i++)
	        	{
	        		QuestList[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}