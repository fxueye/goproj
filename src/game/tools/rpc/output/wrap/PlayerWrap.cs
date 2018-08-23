using System;
using Common.Net.Simple;

namespace Game 
{
	public class PlayerWrap : IWrapper
	{	
		public long UID; // 游戏的uid
		public string GUID; // 平台唯一ID,游客登录的话为设备ID
		public string PID; // 登录平台ID
		public string UserName; // 用户名（唯一，初始和UID一致）
		public long LoginTime; // 登录时间戳（秒）
		public long OnlineTime; // 总在线时长（秒）
		public long CreateTime; // 创角时间戳（秒）
		public long ServerTime; // 服务器时间（秒）
		public int OnBattleIdx; // 上阵索引
		public ItemDataWrap[] Items; // 道具列表
		public int Icon; // 头像
		public long TutorialMask; // 新手引导
		
		public void Decode(Packet pck)
		{	
			UID = pck.GetLong(); 
			GUID = pck.GetString(); 
			PID = pck.GetString(); 
			UserName = pck.GetString(); 
			LoginTime = pck.GetLong(); 
			OnlineTime = pck.GetLong(); 
			CreateTime = pck.GetLong(); 
			ServerTime = pck.GetLong(); 
			OnBattleIdx = pck.GetInt(); 
			Items = new ItemDataWrap[pck.GetShort()];
			for (int i = 0; i < Items.Length; i++)
			{
				Items[i] = new ItemDataWrap();
				Items[i].Decode(pck);
				
			}
			
			Icon = pck.GetInt(); 
			TutorialMask = pck.GetLong(); 
		}
		
        public void Encode(Packet pck)
        {	
        	pck.PutLong(UID); 
        	pck.PutString(GUID); 
        	pck.PutString(PID); 
        	pck.PutString(UserName); 
        	pck.PutLong(LoginTime); 
        	pck.PutLong(OnlineTime); 
        	pck.PutLong(CreateTime); 
        	pck.PutLong(ServerTime); 
        	pck.PutInt(OnBattleIdx); 
        	if (Items == null) pck.PutShort((short)0); 
        	else
        	{
	        	pck.PutShort((short)Items.Length);
	        	for(int i = 0; i < Items.Length; i++)
	        	{
	        		Items[i].Encode(pck);
					
	        	}
	        }
        	
        	pck.PutInt(Icon); 
        	pck.PutLong(TutorialMask); 
        }
	}
}