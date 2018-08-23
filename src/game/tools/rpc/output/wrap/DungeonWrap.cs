using System;
using Common.Net.Simple;

namespace Game 
{
	public class DungeonWrap : IWrapper
	{	
		public int ID; // 活动本id
		public short OpenCnt; // 当前开启次数
		public int RestoreCount; // 当前续命次数
		public int HP; // 当前血量
		public int MaxHP; // 最大血量
		public int Step; // 当前关卡
		public int[] DropHero; // 掉落英雄
		public int[] DropCount; // 掉落数量（与掉落英雄对应）
		
		public void Decode(Packet pck)
		{	
			ID = pck.GetInt(); 
			OpenCnt = pck.GetShort(); 
			RestoreCount = pck.GetInt(); 
			HP = pck.GetInt(); 
			MaxHP = pck.GetInt(); 
			Step = pck.GetInt(); 
			DropHero = new int[pck.GetShort()];
			for (int i = 0; i < DropHero.Length; i++)
			{
				DropHero[i] = pck.GetInt();
			}
			
			DropCount = new int[pck.GetShort()];
			for (int i = 0; i < DropCount.Length; i++)
			{
				DropCount[i] = pck.GetInt();
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(ID); 
        	pck.PutShort(OpenCnt); 
        	pck.PutInt(RestoreCount); 
        	pck.PutInt(HP); 
        	pck.PutInt(MaxHP); 
        	pck.PutInt(Step); 
        	if (DropHero == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)DropHero.Length);
	        	for(int i = 0; i < DropHero.Length; i++)
	        	{
	        		pck.PutInt(DropHero[i]);
	        	}
	        }
        	
        	if (DropCount == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)DropCount.Length);
	        	for(int i = 0; i < DropCount.Length; i++)
	        	{
	        		pck.PutInt(DropCount[i]);
	        	}
	        }
        	
        }
	}
}