using System;
using Common.Log;
using Common.Net.Simple;

namespace Game
{
	public interface IClientCmds 
	{
		void HeartBeat(Command cmd); // 心跳
		void LoginSuccess(Command cmd, PlayerWrap player, bool reconnect, string extension); // 登录成功
		void LoginFailed(Command cmd, short errorCode, string errMsg); // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
		void BindResult(Command cmd, short errCode, string GUID, string pID); // 绑定结果 errCode 0:成功 1：此UID已经绑过guid 2：此GUID已经被绑过，其他：未知错误
		void PlayerError(Command cmd, int errorCode); // 错误码见（ErrorConfig)
		void GameVersionError(Command cmd, string serverGameVersion); // 版本校验失败
		void GSInfo(Command cmd, string ip, int port, string token); // GS服务器地址
		void LoginGSResult(Command cmd, int errorCode); // 登录GS结果
		void SetUserNameAck(Command cmd, short errorCode, string newName); // 改名结果 errcode(-1 内部错误，0-成功，1-重名，2-冷却，3-非法字符)
		void SetIconAck(Command cmd, short errorCode, int iconID); // 修改头像 errcode(0-成功,1-失败)
		void BattleStartQuest(Command cmd, int dungeonID, int questID, long seed, long timestamp, BattlePlayerWrap battlePlayer); // 开始战斗（副本）
		void BattleStartPVP(Command cmd, short battleType, long seed, long timestamp, short firstTurnCamp, BattlePlayerWrap playerA, BattlePlayerWrap playerB); // 开始战斗（竞技）
		void BattleStartReport(Command cmd, BattleReportWrap report); // 战斗开始（战报）
		void BattleCmd(Command cmd, BattleCmdWrap battleCmd); // 战斗指令
		void BattleCmdAck(Command cmd, int seqID); // 战斗指令确认
		void PvPMatching(Command cmd, short pvpType); // pvp匹配中
		void PvPCanceled(Command cmd, short pvpType); // pvp取消
		void ArenaData(Command cmd, ArenaWrap data); // 竞技场随机英雄
		void LadderSeasonData(Command cmd, int seasonID, long seasonEndTime); // 天梯赛季数据
		void LadderTopList(Command cmd, TopListWrap list); // 
		void HerosList(Command cmd, HeroListDataWrap data); // 英雄列表
		void AddHero(Command cmd, short code, int heroID, int count); // 添加英雄卡牌， code（0 未知，1 升级，2 抽卡）
		void UpdateHero(Command cmd, short code, HeroDataWrap data); // 更新英雄卡牌， code（0 未知，1 升级，2 抽卡）
		void AddItem(Command cmd, short code, int itemID, int count); // 添加道具信息， code（0 未知）
		void UpdateItem(Command cmd, short code, ItemDataWrap data); // 更新道具信息， code（0 未知）
		void ChestList(Command cmd, ChestListDataWrap data); // 宝箱列表数据
		void AddOrUpdateChest(Command cmd, short code, ChestDataWrap data); // 更新宝箱信息， code（0 未知）
		void OpenChest(Command cmd, short code, ChestOpenDataWrap data); // 开启宝箱， code（0 未知）
		void QuestList(Command cmd, QuestDataListWrap data); // 玩家关卡数据列表
		void QuestUpdate(Command cmd, QuestDataWrap data); // 关卡数据更新
		void DailyGiftList(Command cmd, DailyGiftListDataWrap data); // 每日礼物数据列表
		void UpdateDailyGift(Command cmd, DailyGiftUpdateDataWrap data); // 领取每次礼物成功
		void UpdateDailyActivity(Command cmd, DailyGetActivityDataWrap data); // 领取活跃度宝箱成功
		void DungeonInfo(Command cmd, int groupID, long endTime); // 副本信息
		void MyDungeonData(Command cmd, DungeonListWrap data); // 我的副本数据
		void ActivityDungeonInfo(Command cmd, ActivityDungeonListWrap data); // 活动副本信息
		void ShopItemList(Command cmd, ShopItemListDataWrap data); // 商品数据列表
		void OpenShopChest(Command cmd, ChestOpenDataWrap chestdata, ShopItemDataWrap data); // 跟新商品宝箱数据
		void GetShopHero(Command cmd, HeroListDataWrap herodatas, ShopItemDataWrap data); // 跟新商品英雄数据
		void PayOrderInfo(Command cmd, string orderID, string productID, string extension); // 充值订单信息
		void GuildList(Command cmd, GuildListWrap data); // 公会列表
		void GuildDetail(Command cmd, GuildDetailWrap data); // 查询公会详情
		void MyGuildData(Command cmd, GuildDetailWrap data); // 我的公会信息
		void GuildJoinRequestList(Command cmd, GuildJoinRequestListWrap data); // 入会申请列表
		void GuildJoinRequestAdd(Command cmd, GuildJoinRequestWrap data); // 新增一条入会申请
		void GuildMemberData(Command cmd, GuildMemberBriefWrap data); // 新会员加入公会
		void GuildMemberQuit(Command cmd, GuildMemberBriefWrap data, short reason); // 会员离开公会
		void GuildChatList(Command cmd, GuildChatListWrap data); // 公会聊天列表
		void GuildNewChat(Command cmd, GuildChatWrap data); // 公会聊天消息
		void GuildDemandList(Command cmd, GuildDemandListWrap data); // 公会捐献请求列表
		void GuildDemand(Command cmd, short flag, GuildDemandWrap data); // 公会捐献请求 flag: 0新增1更新2-删除
		void GuildPvPRequest(Command cmd, GuildPvPRequestWrap data); // 公会友谊赛请求
		void GuildPvPRequestDelete(Command cmd, long memberID); // 公会友谊赛请求删除
		void GuildPvPResult(Command cmd, string reportID, BattleReportWrap report); // 公会友谊赛结果
		void GuildLogList(Command cmd, GuildLogListWrap data); // 公会日志列表
		void GuildLog(Command cmd, GuildLogWrap data); // 公会日志
		void GuildMemberList(Command cmd, GuildMemberListWrap data); // 公会成员列表
		void GuildPvPRequestSuccess(Command cmd); // 公会友谊赛请求成功
		void GuildPvPCancelSuccess(Command cmd); // 公会友谊赛取消成功
		void GuildTopList(Command cmd, GuildListWrap data); // 公会排行榜
		void GuildMemberDetail(Command cmd, PlayerDetailWrap data); // 公会成员详情
		void BattleReportList(Command cmd, BattleReportListWrap data); // 个人战报列表
		void BattleReport(Command cmd, BattleReportWrap report); // 个人战报
		void UpdateEmailsData(Command cmd, int operationCode, EmailListWrap data); // 更新用户邮件数据 0表示更新 1表示删除
		void SetReaderEmailsData(Command cmd, EmailListWrap data); // 邮件设置已读
		void GetUserEmailsData(Command cmd, EmailListWrap data); // 获取用户邮件
		void GetAchievementsListData(Command cmd, AchievementListWrap list); // 用户成就数据
		void UpdateAchievementData(Command cmd, AchievementWrap data); // 更新用户成就数据
		void Emotion(Command cmd, int campID, int emotionID); // 发送表情
		void GetCommentListData(Command cmd, HeroCommentListDataWrap list); // 用户评论数据
		void CommentSuccess(Command cmd, HeroCommentDataWrap data); // 用户评论成功
		void CommentLikeSuccess(Command cmd, int hid, int cid); // 用户评论点赞成功
		void LastBattleResult(Command cmd, short result); // 最后一次战斗结果
		
	}
	
    public class ClientCmdsInvoker : IInvoker
    {
        IClientCmds _cmds = null;
		Action<short> _onCmdInvoked = null;
        public ClientCmdsInvoker(IClientCmds cmds)
        {
            _cmds = cmds;
        }
		public void SetOnCmdInvoked(Action<short> onCmdInvoked)
		{
			_onCmdInvoked = onCmdInvoked;
		}

        public void Invoke(Command cmd)
        {
        	try
        	{
        		Packet pack = cmd.Pack;
	        	switch (cmd.Opcode)
	            {
	            	case (short)0: _cmds.HeartBeat(cmd); break;
	            	case (short)1: _cmds.LoginSuccess(cmd, (PlayerWrap)PackUtil.Unpack(typeof(PlayerWrap), pack), pack.GetBool(), pack.GetString()); break;
	            	case (short)2: _cmds.LoginFailed(cmd, pack.GetShort(), pack.GetString()); break;
	            	case (short)3: _cmds.BindResult(cmd, pack.GetShort(), pack.GetString(), pack.GetString()); break;
	            	case (short)4: _cmds.PlayerError(cmd, pack.GetInt()); break;
	            	case (short)5: _cmds.GameVersionError(cmd, pack.GetString()); break;
	            	case (short)6: _cmds.GSInfo(cmd, pack.GetString(), pack.GetInt(), pack.GetString()); break;
	            	case (short)7: _cmds.LoginGSResult(cmd, pack.GetInt()); break;
	            	case (short)8: _cmds.SetUserNameAck(cmd, pack.GetShort(), pack.GetString()); break;
	            	case (short)9: _cmds.SetIconAck(cmd, pack.GetShort(), pack.GetInt()); break;
	            	case (short)10: _cmds.BattleStartQuest(cmd, pack.GetInt(), pack.GetInt(), pack.GetLong(), pack.GetLong(), (BattlePlayerWrap)PackUtil.Unpack(typeof(BattlePlayerWrap), pack)); break;
	            	case (short)11: _cmds.BattleStartPVP(cmd, pack.GetShort(), pack.GetLong(), pack.GetLong(), pack.GetShort(), (BattlePlayerWrap)PackUtil.Unpack(typeof(BattlePlayerWrap), pack), (BattlePlayerWrap)PackUtil.Unpack(typeof(BattlePlayerWrap), pack)); break;
	            	case (short)12: _cmds.BattleStartReport(cmd, (BattleReportWrap)PackUtil.Unpack(typeof(BattleReportWrap), pack)); break;
	            	case (short)13: _cmds.BattleCmd(cmd, (BattleCmdWrap)PackUtil.Unpack(typeof(BattleCmdWrap), pack)); break;
	            	case (short)14: _cmds.BattleCmdAck(cmd, pack.GetInt()); break;
	            	case (short)15: _cmds.PvPMatching(cmd, pack.GetShort()); break;
	            	case (short)16: _cmds.PvPCanceled(cmd, pack.GetShort()); break;
	            	case (short)17: _cmds.ArenaData(cmd, (ArenaWrap)PackUtil.Unpack(typeof(ArenaWrap), pack)); break;
	            	case (short)18: _cmds.LadderSeasonData(cmd, pack.GetInt(), pack.GetLong()); break;
	            	case (short)19: _cmds.LadderTopList(cmd, (TopListWrap)PackUtil.Unpack(typeof(TopListWrap), pack)); break;
	            	case (short)20: _cmds.HerosList(cmd, (HeroListDataWrap)PackUtil.Unpack(typeof(HeroListDataWrap), pack)); break;
	            	case (short)21: _cmds.AddHero(cmd, pack.GetShort(), pack.GetInt(), pack.GetInt()); break;
	            	case (short)22: _cmds.UpdateHero(cmd, pack.GetShort(), (HeroDataWrap)PackUtil.Unpack(typeof(HeroDataWrap), pack)); break;
	            	case (short)23: _cmds.AddItem(cmd, pack.GetShort(), pack.GetInt(), pack.GetInt()); break;
	            	case (short)24: _cmds.UpdateItem(cmd, pack.GetShort(), (ItemDataWrap)PackUtil.Unpack(typeof(ItemDataWrap), pack)); break;
	            	case (short)30: _cmds.ChestList(cmd, (ChestListDataWrap)PackUtil.Unpack(typeof(ChestListDataWrap), pack)); break;
	            	case (short)31: _cmds.AddOrUpdateChest(cmd, pack.GetShort(), (ChestDataWrap)PackUtil.Unpack(typeof(ChestDataWrap), pack)); break;
	            	case (short)32: _cmds.OpenChest(cmd, pack.GetShort(), (ChestOpenDataWrap)PackUtil.Unpack(typeof(ChestOpenDataWrap), pack)); break;
	            	case (short)40: _cmds.QuestList(cmd, (QuestDataListWrap)PackUtil.Unpack(typeof(QuestDataListWrap), pack)); break;
	            	case (short)41: _cmds.QuestUpdate(cmd, (QuestDataWrap)PackUtil.Unpack(typeof(QuestDataWrap), pack)); break;
	            	case (short)50: _cmds.DailyGiftList(cmd, (DailyGiftListDataWrap)PackUtil.Unpack(typeof(DailyGiftListDataWrap), pack)); break;
	            	case (short)51: _cmds.UpdateDailyGift(cmd, (DailyGiftUpdateDataWrap)PackUtil.Unpack(typeof(DailyGiftUpdateDataWrap), pack)); break;
	            	case (short)52: _cmds.UpdateDailyActivity(cmd, (DailyGetActivityDataWrap)PackUtil.Unpack(typeof(DailyGetActivityDataWrap), pack)); break;
	            	case (short)60: _cmds.DungeonInfo(cmd, pack.GetInt(), pack.GetLong()); break;
	            	case (short)61: _cmds.MyDungeonData(cmd, (DungeonListWrap)PackUtil.Unpack(typeof(DungeonListWrap), pack)); break;
	            	case (short)62: _cmds.ActivityDungeonInfo(cmd, (ActivityDungeonListWrap)PackUtil.Unpack(typeof(ActivityDungeonListWrap), pack)); break;
	            	case (short)70: _cmds.ShopItemList(cmd, (ShopItemListDataWrap)PackUtil.Unpack(typeof(ShopItemListDataWrap), pack)); break;
	            	case (short)71: _cmds.OpenShopChest(cmd, (ChestOpenDataWrap)PackUtil.Unpack(typeof(ChestOpenDataWrap), pack), (ShopItemDataWrap)PackUtil.Unpack(typeof(ShopItemDataWrap), pack)); break;
	            	case (short)72: _cmds.GetShopHero(cmd, (HeroListDataWrap)PackUtil.Unpack(typeof(HeroListDataWrap), pack), (ShopItemDataWrap)PackUtil.Unpack(typeof(ShopItemDataWrap), pack)); break;
	            	case (short)73: _cmds.PayOrderInfo(cmd, pack.GetString(), pack.GetString(), pack.GetString()); break;
	            	case (short)80: _cmds.GuildList(cmd, (GuildListWrap)PackUtil.Unpack(typeof(GuildListWrap), pack)); break;
	            	case (short)81: _cmds.GuildDetail(cmd, (GuildDetailWrap)PackUtil.Unpack(typeof(GuildDetailWrap), pack)); break;
	            	case (short)82: _cmds.MyGuildData(cmd, (GuildDetailWrap)PackUtil.Unpack(typeof(GuildDetailWrap), pack)); break;
	            	case (short)83: _cmds.GuildJoinRequestList(cmd, (GuildJoinRequestListWrap)PackUtil.Unpack(typeof(GuildJoinRequestListWrap), pack)); break;
	            	case (short)84: _cmds.GuildJoinRequestAdd(cmd, (GuildJoinRequestWrap)PackUtil.Unpack(typeof(GuildJoinRequestWrap), pack)); break;
	            	case (short)85: _cmds.GuildMemberData(cmd, (GuildMemberBriefWrap)PackUtil.Unpack(typeof(GuildMemberBriefWrap), pack)); break;
	            	case (short)86: _cmds.GuildMemberQuit(cmd, (GuildMemberBriefWrap)PackUtil.Unpack(typeof(GuildMemberBriefWrap), pack), pack.GetShort()); break;
	            	case (short)87: _cmds.GuildChatList(cmd, (GuildChatListWrap)PackUtil.Unpack(typeof(GuildChatListWrap), pack)); break;
	            	case (short)88: _cmds.GuildNewChat(cmd, (GuildChatWrap)PackUtil.Unpack(typeof(GuildChatWrap), pack)); break;
	            	case (short)89: _cmds.GuildDemandList(cmd, (GuildDemandListWrap)PackUtil.Unpack(typeof(GuildDemandListWrap), pack)); break;
	            	case (short)90: _cmds.GuildDemand(cmd, pack.GetShort(), (GuildDemandWrap)PackUtil.Unpack(typeof(GuildDemandWrap), pack)); break;
	            	case (short)91: _cmds.GuildPvPRequest(cmd, (GuildPvPRequestWrap)PackUtil.Unpack(typeof(GuildPvPRequestWrap), pack)); break;
	            	case (short)92: _cmds.GuildPvPRequestDelete(cmd, pack.GetLong()); break;
	            	case (short)93: _cmds.GuildPvPResult(cmd, pack.GetString(), (BattleReportWrap)PackUtil.Unpack(typeof(BattleReportWrap), pack)); break;
	            	case (short)94: _cmds.GuildLogList(cmd, (GuildLogListWrap)PackUtil.Unpack(typeof(GuildLogListWrap), pack)); break;
	            	case (short)95: _cmds.GuildLog(cmd, (GuildLogWrap)PackUtil.Unpack(typeof(GuildLogWrap), pack)); break;
	            	case (short)96: _cmds.GuildMemberList(cmd, (GuildMemberListWrap)PackUtil.Unpack(typeof(GuildMemberListWrap), pack)); break;
	            	case (short)97: _cmds.GuildPvPRequestSuccess(cmd); break;
	            	case (short)98: _cmds.GuildPvPCancelSuccess(cmd); break;
	            	case (short)99: _cmds.GuildTopList(cmd, (GuildListWrap)PackUtil.Unpack(typeof(GuildListWrap), pack)); break;
	            	case (short)100: _cmds.GuildMemberDetail(cmd, (PlayerDetailWrap)PackUtil.Unpack(typeof(PlayerDetailWrap), pack)); break;
	            	case (short)101: _cmds.BattleReportList(cmd, (BattleReportListWrap)PackUtil.Unpack(typeof(BattleReportListWrap), pack)); break;
	            	case (short)102: _cmds.BattleReport(cmd, (BattleReportWrap)PackUtil.Unpack(typeof(BattleReportWrap), pack)); break;
	            	case (short)110: _cmds.UpdateEmailsData(cmd, pack.GetInt(), (EmailListWrap)PackUtil.Unpack(typeof(EmailListWrap), pack)); break;
	            	case (short)111: _cmds.SetReaderEmailsData(cmd, (EmailListWrap)PackUtil.Unpack(typeof(EmailListWrap), pack)); break;
	            	case (short)112: _cmds.GetUserEmailsData(cmd, (EmailListWrap)PackUtil.Unpack(typeof(EmailListWrap), pack)); break;
	            	case (short)113: _cmds.GetAchievementsListData(cmd, (AchievementListWrap)PackUtil.Unpack(typeof(AchievementListWrap), pack)); break;
	            	case (short)114: _cmds.UpdateAchievementData(cmd, (AchievementWrap)PackUtil.Unpack(typeof(AchievementWrap), pack)); break;
	            	case (short)115: _cmds.Emotion(cmd, pack.GetInt(), pack.GetInt()); break;
	            	case (short)120: _cmds.GetCommentListData(cmd, (HeroCommentListDataWrap)PackUtil.Unpack(typeof(HeroCommentListDataWrap), pack)); break;
	            	case (short)121: _cmds.CommentSuccess(cmd, (HeroCommentDataWrap)PackUtil.Unpack(typeof(HeroCommentDataWrap), pack)); break;
	            	case (short)122: _cmds.CommentLikeSuccess(cmd, pack.GetInt(), pack.GetInt()); break;
	            	case (short)123: _cmds.LastBattleResult(cmd, pack.GetShort()); break;
	            	
	            }
				if (_onCmdInvoked != null)
					_onCmdInvoked(cmd.Opcode);
        	}
        	catch(Exception e)
        	{
				L.Error("invoke error, opcode=" + cmd.Opcode);
        		L.Exception(e.Message, e);
        	}
		}
    }
}
