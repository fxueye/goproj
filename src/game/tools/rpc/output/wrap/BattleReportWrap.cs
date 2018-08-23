using System;
using Common.Net.Simple;

namespace Game 
{
	public class BattleReportWrap : IWrapper
	{	
		public int Version; // 战报版本号
		public string ReportID; // 战报ID，服务端生成
		public short BattleType; // 战斗类型（0 天梯，1 竞技场，2 练习赛，3 关卡）
		public int DungeonID; // 活动副本ID(战斗类型为关卡时使用）
		public int QuestID; // 关卡ID(战斗类型为关卡时使用）
		public long Seed; // 随机数种子
		public long Timestamp; // 时间戳
		public short FirstCamp; // 先手阵营
		public short Result; // 战斗结果(1 A赢，2 B赢，3 平局，4 A掉线，5 B掉线）
		public BattlePlayerWrap CampA; // 阵营A
		public BattlePlayerWrap CampB; // 阵营B（战斗类型不为关卡时使用）
		public BattleCmdWrap[] Cmds; // 指令列表
		
		public void Decode(Packet pck)
		{	
			Version = pck.GetInt(); 
			ReportID = pck.GetString(); 
			BattleType = pck.GetShort(); 
			DungeonID = pck.GetInt(); 
			QuestID = pck.GetInt(); 
			Seed = pck.GetLong(); 
			Timestamp = pck.GetLong(); 
			FirstCamp = pck.GetShort(); 
			Result = pck.GetShort(); 
			CampA = new BattlePlayerWrap();
			CampA.Decode(pck);
			
			CampB = new BattlePlayerWrap();
			CampB.Decode(pck);
			
			Cmds = new BattleCmdWrap[pck.GetShort()];
			for (int i = 0; i < Cmds.Length; i++)
			{
				Cmds[i] = new BattleCmdWrap();
				Cmds[i].Decode(pck);
				
			}
			
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutInt(Version); 
        	pck.PutString(ReportID); 
        	pck.PutShort(BattleType); 
        	pck.PutInt(DungeonID); 
        	pck.PutInt(QuestID); 
        	pck.PutLong(Seed); 
        	pck.PutLong(Timestamp); 
        	pck.PutShort(FirstCamp); 
        	pck.PutShort(Result); 
        	CampA.Encode(pck); 
        	CampB.Encode(pck); 
        	if (Cmds == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Cmds.Length);
	        	for(int i = 0; i < Cmds.Length; i++)
	        	{
	        		Cmds[i].Encode(pck);
					
	        	}
	        }
        	
        }
	}
}