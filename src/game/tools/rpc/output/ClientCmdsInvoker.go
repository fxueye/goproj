package cmds
import (
	"errors"
	"fmt"
	rpc "tipcat.com/common/rpc/simple"
	tcp "tipcat.com/common/server/tcp"
	wraps "tipcat.com/cmds/wraps"
)

type IClientCmds interface {
	HeartBeat(cmd *rpc.SimpleCmd, se *tcp.Session) // 心跳
	LoginSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, player *wraps.PlayerWrap, reconnect bool, extension string) // 登录成功
	LoginFailed(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int16, errMsg string) // 登录失败（1 用户不存在，2  密码错误， 3 禁止登陆）
	BindResult(cmd *rpc.SimpleCmd, se *tcp.Session, errCode int16, GUID string, pID string) // 绑定结果 errCode 0:成功 1：此UID已经绑过guid 2：此GUID已经被绑过，其他：未知错误
	PlayerError(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int32) // 错误码见（ErrorConfig)
	GameVersionError(cmd *rpc.SimpleCmd, se *tcp.Session, serverGameVersion string) // 版本校验失败
	GSInfo(cmd *rpc.SimpleCmd, se *tcp.Session, ip string, port int32, token string) // GS服务器地址
	LoginGSResult(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int32) // 登录GS结果
	SetUserNameAck(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int16, newName string) // 改名结果 errcode(-1 内部错误，0-成功，1-重名，2-冷却，3-非法字符)
	SetIconAck(cmd *rpc.SimpleCmd, se *tcp.Session, errorCode int16, iconID int32) // 修改头像 errcode(0-成功,1-失败)
	BattleStartQuest(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32, questID int32, seed int64, timestamp int64, battlePlayer *wraps.BattlePlayerWrap) // 开始战斗（副本）
	BattleStartPVP(cmd *rpc.SimpleCmd, se *tcp.Session, battleType int16, seed int64, timestamp int64, firstTurnCamp int16, playerA *wraps.BattlePlayerWrap, playerB *wraps.BattlePlayerWrap) // 开始战斗（竞技）
	BattleStartReport(cmd *rpc.SimpleCmd, se *tcp.Session, report *wraps.BattleReportWrap) // 战斗开始（战报）
	BattleCmd(cmd *rpc.SimpleCmd, se *tcp.Session, battleCmd *wraps.BattleCmdWrap) // 战斗指令
	BattleCmdAck(cmd *rpc.SimpleCmd, se *tcp.Session, seqID int32) // 战斗指令确认
	PvPMatching(cmd *rpc.SimpleCmd, se *tcp.Session, pvpType int16) // pvp匹配中
	PvPCanceled(cmd *rpc.SimpleCmd, se *tcp.Session, pvpType int16) // pvp取消
	ArenaData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.ArenaWrap) // 竞技场随机英雄
	LadderSeasonData(cmd *rpc.SimpleCmd, se *tcp.Session, seasonID int32, seasonEndTime int64) // 天梯赛季数据
	LadderTopList(cmd *rpc.SimpleCmd, se *tcp.Session, list *wraps.TopListWrap) // 
	HerosList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.HeroListDataWrap) // 英雄列表
	AddHero(cmd *rpc.SimpleCmd, se *tcp.Session, code int16, heroID int32, count int32) // 添加英雄卡牌， code（0 未知，1 升级，2 抽卡）
	UpdateHero(cmd *rpc.SimpleCmd, se *tcp.Session, code int16, data *wraps.HeroDataWrap) // 更新英雄卡牌， code（0 未知，1 升级，2 抽卡）
	AddItem(cmd *rpc.SimpleCmd, se *tcp.Session, code int16, itemID int32, count int32) // 添加道具信息， code（0 未知）
	UpdateItem(cmd *rpc.SimpleCmd, se *tcp.Session, code int16, data *wraps.ItemDataWrap) // 更新道具信息， code（0 未知）
	ChestList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.ChestListDataWrap) // 宝箱列表数据
	AddOrUpdateChest(cmd *rpc.SimpleCmd, se *tcp.Session, code int16, data *wraps.ChestDataWrap) // 更新宝箱信息， code（0 未知）
	OpenChest(cmd *rpc.SimpleCmd, se *tcp.Session, code int16, data *wraps.ChestOpenDataWrap) // 开启宝箱， code（0 未知）
	QuestList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.QuestDataListWrap) // 玩家关卡数据列表
	QuestUpdate(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.QuestDataWrap) // 关卡数据更新
	DailyGiftList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.DailyGiftListDataWrap) // 每日礼物数据列表
	UpdateDailyGift(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.DailyGiftUpdateDataWrap) // 领取每次礼物成功
	UpdateDailyActivity(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.DailyGetActivityDataWrap) // 领取活跃度宝箱成功
	DungeonInfo(cmd *rpc.SimpleCmd, se *tcp.Session, groupID int32, endTime int64) // 副本信息
	MyDungeonData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.DungeonListWrap) // 我的副本数据
	ActivityDungeonInfo(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.ActivityDungeonListWrap) // 活动副本信息
	ShopItemList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.ShopItemListDataWrap) // 商品数据列表
	OpenShopChest(cmd *rpc.SimpleCmd, se *tcp.Session, chestdata *wraps.ChestOpenDataWrap, data *wraps.ShopItemDataWrap) // 跟新商品宝箱数据
	GetShopHero(cmd *rpc.SimpleCmd, se *tcp.Session, herodatas *wraps.HeroListDataWrap, data *wraps.ShopItemDataWrap) // 跟新商品英雄数据
	PayOrderInfo(cmd *rpc.SimpleCmd, se *tcp.Session, orderID string, productID string, extension string) // 充值订单信息
	GuildList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildListWrap) // 公会列表
	GuildDetail(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildDetailWrap) // 查询公会详情
	MyGuildData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildDetailWrap) // 我的公会信息
	GuildJoinRequestList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildJoinRequestListWrap) // 入会申请列表
	GuildJoinRequestAdd(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildJoinRequestWrap) // 新增一条入会申请
	GuildMemberData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildMemberBriefWrap) // 新会员加入公会
	GuildMemberQuit(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildMemberBriefWrap, reason int16) // 会员离开公会
	GuildChatList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildChatListWrap) // 公会聊天列表
	GuildNewChat(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildChatWrap) // 公会聊天消息
	GuildDemandList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildDemandListWrap) // 公会捐献请求列表
	GuildDemand(cmd *rpc.SimpleCmd, se *tcp.Session, flag int16, data *wraps.GuildDemandWrap) // 公会捐献请求 flag: 0新增1更新2-删除
	GuildPvPRequest(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildPvPRequestWrap) // 公会友谊赛请求
	GuildPvPRequestDelete(cmd *rpc.SimpleCmd, se *tcp.Session, memberID int64) // 公会友谊赛请求删除
	GuildPvPResult(cmd *rpc.SimpleCmd, se *tcp.Session, reportID string, report *wraps.BattleReportWrap) // 公会友谊赛结果
	GuildLogList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildLogListWrap) // 公会日志列表
	GuildLog(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildLogWrap) // 公会日志
	GuildMemberList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildMemberListWrap) // 公会成员列表
	GuildPvPRequestSuccess(cmd *rpc.SimpleCmd, se *tcp.Session) // 公会友谊赛请求成功
	GuildPvPCancelSuccess(cmd *rpc.SimpleCmd, se *tcp.Session) // 公会友谊赛取消成功
	GuildTopList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildListWrap) // 公会排行榜
	GuildMemberDetail(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.PlayerDetailWrap) // 公会成员详情
	BattleReportList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.BattleReportListWrap) // 个人战报列表
	BattleReport(cmd *rpc.SimpleCmd, se *tcp.Session, report *wraps.BattleReportWrap) // 个人战报
	UpdateEmailsData(cmd *rpc.SimpleCmd, se *tcp.Session, operationCode int32, data *wraps.EmailListWrap) // 更新用户邮件数据 0表示更新 1表示删除
	SetReaderEmailsData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.EmailListWrap) // 邮件设置已读
	GetUserEmailsData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.EmailListWrap) // 获取用户邮件
	GetAchievementsListData(cmd *rpc.SimpleCmd, se *tcp.Session, list *wraps.AchievementListWrap) // 用户成就数据
	UpdateAchievementData(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.AchievementWrap) // 更新用户成就数据
	Emotion(cmd *rpc.SimpleCmd, se *tcp.Session, campID int32, emotionID int32) // 发送表情
	GetCommentListData(cmd *rpc.SimpleCmd, se *tcp.Session, list *wraps.HeroCommentListDataWrap) // 用户评论数据
	CommentSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.HeroCommentDataWrap) // 用户评论成功
	CommentLikeSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, hid int32, cid int32) // 用户评论点赞成功
	LastBattleResult(cmd *rpc.SimpleCmd, se *tcp.Session, result int16) // 最后一次战斗结果
	
}

type ClientCmdsInvoker struct {
	invoker IClientCmds
	defaultInvoker func(cmd *rpc.SimpleCmd, se *tcp.Session)
	rpc.SimpleInvoker
} 

func NewClientCmdsInvoker(invoker IClientCmds, defaultInvoker func(*rpc.SimpleCmd, *tcp.Session)) *ClientCmdsInvoker {
	inv := new(ClientCmdsInvoker)
	inv.invoker = invoker
	inv.defaultInvoker = defaultInvoker
	return inv
} 

func (this *ClientCmdsInvoker) Invoke(cmd *rpc.SimpleCmd, se *tcp.Session) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	pack := cmd.Pack()
	switch(cmd.Opcode()) {
	case 0: 
		this.invoker.HeartBeat(cmd,se)
	case 1: 
		this.invoker.LoginSuccess(cmd,se, new(wraps.PlayerWrap).Decode(pack).(*wraps.PlayerWrap), pack.PopBool(), pack.PopString())
	case 2: 
		this.invoker.LoginFailed(cmd,se, pack.PopInt16(), pack.PopString())
	case 3: 
		this.invoker.BindResult(cmd,se, pack.PopInt16(), pack.PopString(), pack.PopString())
	case 4: 
		this.invoker.PlayerError(cmd,se, pack.PopInt32())
	case 5: 
		this.invoker.GameVersionError(cmd,se, pack.PopString())
	case 6: 
		this.invoker.GSInfo(cmd,se, pack.PopString(), pack.PopInt32(), pack.PopString())
	case 7: 
		this.invoker.LoginGSResult(cmd,se, pack.PopInt32())
	case 8: 
		this.invoker.SetUserNameAck(cmd,se, pack.PopInt16(), pack.PopString())
	case 9: 
		this.invoker.SetIconAck(cmd,se, pack.PopInt16(), pack.PopInt32())
	case 10: 
		this.invoker.BattleStartQuest(cmd,se, pack.PopInt32(), pack.PopInt32(), pack.PopInt64(), pack.PopInt64(), new(wraps.BattlePlayerWrap).Decode(pack).(*wraps.BattlePlayerWrap))
	case 11: 
		this.invoker.BattleStartPVP(cmd,se, pack.PopInt16(), pack.PopInt64(), pack.PopInt64(), pack.PopInt16(), new(wraps.BattlePlayerWrap).Decode(pack).(*wraps.BattlePlayerWrap), new(wraps.BattlePlayerWrap).Decode(pack).(*wraps.BattlePlayerWrap))
	case 12: 
		this.invoker.BattleStartReport(cmd,se, new(wraps.BattleReportWrap).Decode(pack).(*wraps.BattleReportWrap))
	case 13: 
		this.invoker.BattleCmd(cmd,se, new(wraps.BattleCmdWrap).Decode(pack).(*wraps.BattleCmdWrap))
	case 14: 
		this.invoker.BattleCmdAck(cmd,se, pack.PopInt32())
	case 15: 
		this.invoker.PvPMatching(cmd,se, pack.PopInt16())
	case 16: 
		this.invoker.PvPCanceled(cmd,se, pack.PopInt16())
	case 17: 
		this.invoker.ArenaData(cmd,se, new(wraps.ArenaWrap).Decode(pack).(*wraps.ArenaWrap))
	case 18: 
		this.invoker.LadderSeasonData(cmd,se, pack.PopInt32(), pack.PopInt64())
	case 19: 
		this.invoker.LadderTopList(cmd,se, new(wraps.TopListWrap).Decode(pack).(*wraps.TopListWrap))
	case 20: 
		this.invoker.HerosList(cmd,se, new(wraps.HeroListDataWrap).Decode(pack).(*wraps.HeroListDataWrap))
	case 21: 
		this.invoker.AddHero(cmd,se, pack.PopInt16(), pack.PopInt32(), pack.PopInt32())
	case 22: 
		this.invoker.UpdateHero(cmd,se, pack.PopInt16(), new(wraps.HeroDataWrap).Decode(pack).(*wraps.HeroDataWrap))
	case 23: 
		this.invoker.AddItem(cmd,se, pack.PopInt16(), pack.PopInt32(), pack.PopInt32())
	case 24: 
		this.invoker.UpdateItem(cmd,se, pack.PopInt16(), new(wraps.ItemDataWrap).Decode(pack).(*wraps.ItemDataWrap))
	case 30: 
		this.invoker.ChestList(cmd,se, new(wraps.ChestListDataWrap).Decode(pack).(*wraps.ChestListDataWrap))
	case 31: 
		this.invoker.AddOrUpdateChest(cmd,se, pack.PopInt16(), new(wraps.ChestDataWrap).Decode(pack).(*wraps.ChestDataWrap))
	case 32: 
		this.invoker.OpenChest(cmd,se, pack.PopInt16(), new(wraps.ChestOpenDataWrap).Decode(pack).(*wraps.ChestOpenDataWrap))
	case 40: 
		this.invoker.QuestList(cmd,se, new(wraps.QuestDataListWrap).Decode(pack).(*wraps.QuestDataListWrap))
	case 41: 
		this.invoker.QuestUpdate(cmd,se, new(wraps.QuestDataWrap).Decode(pack).(*wraps.QuestDataWrap))
	case 50: 
		this.invoker.DailyGiftList(cmd,se, new(wraps.DailyGiftListDataWrap).Decode(pack).(*wraps.DailyGiftListDataWrap))
	case 51: 
		this.invoker.UpdateDailyGift(cmd,se, new(wraps.DailyGiftUpdateDataWrap).Decode(pack).(*wraps.DailyGiftUpdateDataWrap))
	case 52: 
		this.invoker.UpdateDailyActivity(cmd,se, new(wraps.DailyGetActivityDataWrap).Decode(pack).(*wraps.DailyGetActivityDataWrap))
	case 60: 
		this.invoker.DungeonInfo(cmd,se, pack.PopInt32(), pack.PopInt64())
	case 61: 
		this.invoker.MyDungeonData(cmd,se, new(wraps.DungeonListWrap).Decode(pack).(*wraps.DungeonListWrap))
	case 62: 
		this.invoker.ActivityDungeonInfo(cmd,se, new(wraps.ActivityDungeonListWrap).Decode(pack).(*wraps.ActivityDungeonListWrap))
	case 70: 
		this.invoker.ShopItemList(cmd,se, new(wraps.ShopItemListDataWrap).Decode(pack).(*wraps.ShopItemListDataWrap))
	case 71: 
		this.invoker.OpenShopChest(cmd,se, new(wraps.ChestOpenDataWrap).Decode(pack).(*wraps.ChestOpenDataWrap), new(wraps.ShopItemDataWrap).Decode(pack).(*wraps.ShopItemDataWrap))
	case 72: 
		this.invoker.GetShopHero(cmd,se, new(wraps.HeroListDataWrap).Decode(pack).(*wraps.HeroListDataWrap), new(wraps.ShopItemDataWrap).Decode(pack).(*wraps.ShopItemDataWrap))
	case 73: 
		this.invoker.PayOrderInfo(cmd,se, pack.PopString(), pack.PopString(), pack.PopString())
	case 80: 
		this.invoker.GuildList(cmd,se, new(wraps.GuildListWrap).Decode(pack).(*wraps.GuildListWrap))
	case 81: 
		this.invoker.GuildDetail(cmd,se, new(wraps.GuildDetailWrap).Decode(pack).(*wraps.GuildDetailWrap))
	case 82: 
		this.invoker.MyGuildData(cmd,se, new(wraps.GuildDetailWrap).Decode(pack).(*wraps.GuildDetailWrap))
	case 83: 
		this.invoker.GuildJoinRequestList(cmd,se, new(wraps.GuildJoinRequestListWrap).Decode(pack).(*wraps.GuildJoinRequestListWrap))
	case 84: 
		this.invoker.GuildJoinRequestAdd(cmd,se, new(wraps.GuildJoinRequestWrap).Decode(pack).(*wraps.GuildJoinRequestWrap))
	case 85: 
		this.invoker.GuildMemberData(cmd,se, new(wraps.GuildMemberBriefWrap).Decode(pack).(*wraps.GuildMemberBriefWrap))
	case 86: 
		this.invoker.GuildMemberQuit(cmd,se, new(wraps.GuildMemberBriefWrap).Decode(pack).(*wraps.GuildMemberBriefWrap), pack.PopInt16())
	case 87: 
		this.invoker.GuildChatList(cmd,se, new(wraps.GuildChatListWrap).Decode(pack).(*wraps.GuildChatListWrap))
	case 88: 
		this.invoker.GuildNewChat(cmd,se, new(wraps.GuildChatWrap).Decode(pack).(*wraps.GuildChatWrap))
	case 89: 
		this.invoker.GuildDemandList(cmd,se, new(wraps.GuildDemandListWrap).Decode(pack).(*wraps.GuildDemandListWrap))
	case 90: 
		this.invoker.GuildDemand(cmd,se, pack.PopInt16(), new(wraps.GuildDemandWrap).Decode(pack).(*wraps.GuildDemandWrap))
	case 91: 
		this.invoker.GuildPvPRequest(cmd,se, new(wraps.GuildPvPRequestWrap).Decode(pack).(*wraps.GuildPvPRequestWrap))
	case 92: 
		this.invoker.GuildPvPRequestDelete(cmd,se, pack.PopInt64())
	case 93: 
		this.invoker.GuildPvPResult(cmd,se, pack.PopString(), new(wraps.BattleReportWrap).Decode(pack).(*wraps.BattleReportWrap))
	case 94: 
		this.invoker.GuildLogList(cmd,se, new(wraps.GuildLogListWrap).Decode(pack).(*wraps.GuildLogListWrap))
	case 95: 
		this.invoker.GuildLog(cmd,se, new(wraps.GuildLogWrap).Decode(pack).(*wraps.GuildLogWrap))
	case 96: 
		this.invoker.GuildMemberList(cmd,se, new(wraps.GuildMemberListWrap).Decode(pack).(*wraps.GuildMemberListWrap))
	case 97: 
		this.invoker.GuildPvPRequestSuccess(cmd,se)
	case 98: 
		this.invoker.GuildPvPCancelSuccess(cmd,se)
	case 99: 
		this.invoker.GuildTopList(cmd,se, new(wraps.GuildListWrap).Decode(pack).(*wraps.GuildListWrap))
	case 100: 
		this.invoker.GuildMemberDetail(cmd,se, new(wraps.PlayerDetailWrap).Decode(pack).(*wraps.PlayerDetailWrap))
	case 101: 
		this.invoker.BattleReportList(cmd,se, new(wraps.BattleReportListWrap).Decode(pack).(*wraps.BattleReportListWrap))
	case 102: 
		this.invoker.BattleReport(cmd,se, new(wraps.BattleReportWrap).Decode(pack).(*wraps.BattleReportWrap))
	case 110: 
		this.invoker.UpdateEmailsData(cmd,se, pack.PopInt32(), new(wraps.EmailListWrap).Decode(pack).(*wraps.EmailListWrap))
	case 111: 
		this.invoker.SetReaderEmailsData(cmd,se, new(wraps.EmailListWrap).Decode(pack).(*wraps.EmailListWrap))
	case 112: 
		this.invoker.GetUserEmailsData(cmd,se, new(wraps.EmailListWrap).Decode(pack).(*wraps.EmailListWrap))
	case 113: 
		this.invoker.GetAchievementsListData(cmd,se, new(wraps.AchievementListWrap).Decode(pack).(*wraps.AchievementListWrap))
	case 114: 
		this.invoker.UpdateAchievementData(cmd,se, new(wraps.AchievementWrap).Decode(pack).(*wraps.AchievementWrap))
	case 115: 
		this.invoker.Emotion(cmd,se, pack.PopInt32(), pack.PopInt32())
	case 120: 
		this.invoker.GetCommentListData(cmd,se, new(wraps.HeroCommentListDataWrap).Decode(pack).(*wraps.HeroCommentListDataWrap))
	case 121: 
		this.invoker.CommentSuccess(cmd,se, new(wraps.HeroCommentDataWrap).Decode(pack).(*wraps.HeroCommentDataWrap))
	case 122: 
		this.invoker.CommentLikeSuccess(cmd,se, pack.PopInt32(), pack.PopInt32())
	case 123: 
		this.invoker.LastBattleResult(cmd,se, pack.PopInt16())
	
	default:
		if this.defaultInvoker != nil {
			this.defaultInvoker(cmd,se)
		}
	}
	return nil
}

