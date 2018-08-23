package cmds
import (
	"errors"
	"fmt"
	rpc "tipcat.com/common/rpc/simple"
	tcp "tipcat.com/common/server/tcp"
	wraps "tipcat.com/cmds/wraps"
)

type IServerCSCmds interface {
	GW2CS_Ping(cmd *rpc.SimpleCmd, se *tcp.Session) // gw心跳
	GS2CS_Ping(cmd *rpc.SimpleCmd, se *tcp.Session) // gs心跳
	GB2CS_Pong(cmd *rpc.SimpleCmd, se *tcp.Session) // gb心跳
	GW2CS_LoginGuest(cmd *rpc.SimpleCmd, se *tcp.Session, deviceID string, deviceType string, partnerID string, ip string) // 登录
	GW2CS_LoginToken(cmd *rpc.SimpleCmd, se *tcp.Session, ptID string, ptData string, deviceType string, partnerID string, ip string, reconnect bool, extension string) // 登录
	GW2CS_BindAccount(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, ptID string, ptData string) // 绑定账户
	GS2CS_Register(cmd *rpc.SimpleCmd, se *tcp.Session, gsIdx int32, ip string, port int32) // 注册
	GW2CS_Logout(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64) // 登出
	GW2CS_SessionDisconnect(cmd *rpc.SimpleCmd, se *tcp.Session, sessionID int64) // session断开
	GB2CS_RechargeNotify(cmd *rpc.SimpleCmd, se *tcp.Session, seqID int32, uid int64, orderID string, payID string, amount float32) // 充值通知
	GB2CS_BanUser(cmd *rpc.SimpleCmd, se *tcp.Session, seqID int32, uid int64, banType int32, banTimestamp int64) // 封停账号
	GB2CS_UnBanUser(cmd *rpc.SimpleCmd, se *tcp.Session, seqID int32, uid int64) // 解封账号
	GB2CS_CreateUserAck(cmd *rpc.SimpleCmd, se *tcp.Session, result int32, guid string, ptid string, loginIP string, deviceType string, partnerId string, ptAcc string, extension string) // global创建角色返回
	GB2CS_UpdateUserScore(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, flag bool, score int32) // 更新用户积分
	GB2CS_DeleteUserData(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, save bool) // 内存删除用户数据
	GB2CS_PvpResult(cmd *rpc.SimpleCmd, se *tcp.Session, report *wraps.BattleReportWrap) // pvp结果
	GB2CS_BattleCombo(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, data *wraps.IntArrayWrap) // 战斗连击
	GB2CS_SeasonTime(cmd *rpc.SimpleCmd, se *tcp.Session, seasonID int32, seasonTime int64) // 赛季时间
	GS2CS_CheckQuestResult(cmd *rpc.SimpleCmd, se *tcp.Session, result bool, report *wraps.BattleReportWrap, hp int32) // 关卡校验结果
	GB2CS_SendSystemEmail(cmd *rpc.SimpleCmd, se *tcp.Session, title string, content string, from string, attachment string, emailID int32) // 后台发送系统邮件
	GB2CS_SendPersonEmail(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, title string, content string, from string, attachment string, wraps *wraps.StringKVArrWrap) // 后台发送个人邮件
	GB2CS_GuildCreateSuccess(cmd *rpc.SimpleCmd, se *tcp.Session) // 公会创建成功
	GB2CS_GuildDemandSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, heroid int32) // 公会索要成功
	GB2CS_GuildDonateSuccess(cmd *rpc.SimpleCmd, se *tcp.Session, heroId int32, count int32) // 公会捐献成功
	GB2CS_SyncTopList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.TopListWrap) // 同步排行榜数据给cs
	GB2CS_SyncMemberList(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.MemberListSyncWrap) // 同步所有公会成员列表
	GB2CS_SyncMember(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.MemberSyncWrap) // 同步公会成员
	GB2CS_ReceiveDonateHero(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64, heroId int32, count int32) // 获得捐献英雄
	GB2CS_SetDungeon(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.ActivityDungeonListWrap) // 设置副本
	Agent2CS_Pong(cmd *rpc.SimpleCmd, se *tcp.Session) // agent心跳回复
	GB2CS_SyncGuild(cmd *rpc.SimpleCmd, se *tcp.Session, data *wraps.GuildSyncWrap) // 同步公会信息
	GW2CS_UpBattleHeros(cmd *rpc.SimpleCmd, se *tcp.Session, battleindex int32, herosdataList *wraps.StringKVArrWrap) // 用户上阵英雄
	GW2CS_ChangeBattleIndex(cmd *rpc.SimpleCmd, se *tcp.Session, battleindex int32) // 用户更换阵容
	GW2CS_HeroLevelup(cmd *rpc.SimpleCmd, se *tcp.Session, heroid int32) // 英雄升级
	GW2CS_SetName(cmd *rpc.SimpleCmd, se *tcp.Session, newName string) // 修改用户名
	GW2CS_SetIcon(cmd *rpc.SimpleCmd, se *tcp.Session, iconID int32) // 设置头像
	GW2CS_PvpMatch(cmd *rpc.SimpleCmd, se *tcp.Session, pvpType int16, battleindex int16) // pvp比赛报名
	GW2CS_PvpCancel(cmd *rpc.SimpleCmd, se *tcp.Session) // pvp比赛取消
	GW2CS_StartQuest(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32, questid int32, battleindex int16) // 开始关卡
	GW2CS_BuyArenaTicket(cmd *rpc.SimpleCmd, se *tcp.Session) // 购买竞技场门票
	GW2CS_ArenaHeroOnBattle(cmd *rpc.SimpleCmd, se *tcp.Session, id1 int32, id2 int32, id3 int32, id4 int32) // 竞技场上阵英雄
	GW2CS_FinishQuest(cmd *rpc.SimpleCmd, se *tcp.Session, report *wraps.BattleReportWrap) // 关卡结束
	GW2CS_GetArenaReward(cmd *rpc.SimpleCmd, se *tcp.Session) // 竞技场领奖
	GW2CS_GetDailyGift(cmd *rpc.SimpleCmd, se *tcp.Session, gidx int32) // 领取每日礼物
	GW2CS_RefreshDailyGift(cmd *rpc.SimpleCmd, se *tcp.Session) // 刷新每日礼物
	GW2CS_GetDailyActivity(cmd *rpc.SimpleCmd, se *tcp.Session) // 获取活跃度宝箱
	GW2CS_ActiveChest(cmd *rpc.SimpleCmd, se *tcp.Session, cidx int32) // 激活宝箱
	GW2CS_OpenChest(cmd *rpc.SimpleCmd, se *tcp.Session, cidx int32) // 开启宝箱
	GW2CS_OpenChestByDiamond(cmd *rpc.SimpleCmd, se *tcp.Session, cidx int32) // 用钻石开启宝箱
	GW2CS_OpenFreeChest(cmd *rpc.SimpleCmd, se *tcp.Session) // 开启免费宝箱
	GW2CS_OpenArenaChest(cmd *rpc.SimpleCmd, se *tcp.Session) // 开启竞技场宝箱
	GW2CS_LadderTopList(cmd *rpc.SimpleCmd, se *tcp.Session, kind int16) // 获取排行榜信息
	GW2CS_OpenDungeon(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32) // 开启活动副本
	GW2CS_RestoreDungeon(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32, restore bool) // 活动副本续命,restore为false代表放弃续命，从头开始
	GW2CS_DungeonLoot(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32) // 获取副本掉落
	GW2CS_DungeonBattleFail(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32) // 活动副本战斗失败
	GW2CS_DungeonAddHp(cmd *rpc.SimpleCmd, se *tcp.Session, dungeonID int32, hpType int16) // 副本补血
	GW2CS_ShopBuyChest(cmd *rpc.SimpleCmd, se *tcp.Session, shopItemId int32) // 商店购买宝箱
	GW2CS_ShopBuyHero(cmd *rpc.SimpleCmd, se *tcp.Session, shopItemId int32) // 商店购买英雄
	GW2CS_ShopBuyGold(cmd *rpc.SimpleCmd, se *tcp.Session, shopItemId int32) // 商店购买金币
	GW2CS_ShopReloadItems(cmd *rpc.SimpleCmd, se *tcp.Session) // 商店刷新物品
	GW2CS_ShopBuyDiamond(cmd *rpc.SimpleCmd, se *tcp.Session, shopItemId int32, extension string) // 商店购买钻石
	GW2CS_CreateGuild(cmd *rpc.SimpleCmd, se *tcp.Session, guildName string, guildNote string, icon int32, joinType int16, joinScore int32) // 创建公会
	GW2CS_QueryGuild(cmd *rpc.SimpleCmd, se *tcp.Session, keyword string, minMember int16, maxMember int16, minScore int32, onlyCanJoin bool) // 查询公会
	GW2CS_QueryGuildDetail(cmd *rpc.SimpleCmd, se *tcp.Session, guildid int32) // 查询公会详情
	GW2CS_RequestJoinGuild(cmd *rpc.SimpleCmd, se *tcp.Session, guildid int32, msg string) // 申请加入公会
	GW2CS_DealJoinRequest(cmd *rpc.SimpleCmd, se *tcp.Session, reqUID int64, agree bool) // 处理入会申请
	GW2CS_LeaveGuild(cmd *rpc.SimpleCmd, se *tcp.Session) // 离开公会
	GW2CS_KickMember(cmd *rpc.SimpleCmd, se *tcp.Session, memberid int64) // 开除会员
	GW2CS_AppointMember(cmd *rpc.SimpleCmd, se *tcp.Session, memberid int64, guildrank int16) // 任免会员
	GW2CS_GuildChat(cmd *rpc.SimpleCmd, se *tcp.Session, msg string) // 公会聊天
	GW2CS_GuildDemand(cmd *rpc.SimpleCmd, se *tcp.Session, heroid int32) // 公会索要英雄
	GW2CS_GuildDonate(cmd *rpc.SimpleCmd, se *tcp.Session, memberid int64, heroId int32, count int32) // 公会捐献
	GW2CS_GuildSetting(cmd *rpc.SimpleCmd, se *tcp.Session, guildNote string, icon int32, joinType int16, joinScore int32) // 公会设置
	GW2CS_GuildRequestPvP(cmd *rpc.SimpleCmd, se *tcp.Session, note string, battleIndex int16) // 发起公会友谊赛
	GW2CS_GuildCancelPvP(cmd *rpc.SimpleCmd, se *tcp.Session) // 取消发起友谊赛
	GW2CS_GuildStartPvP(cmd *rpc.SimpleCmd, se *tcp.Session, memberid int64, battleIndex int16) // 友谊赛应战
	GW2CS_GuildPvPReplay(cmd *rpc.SimpleCmd, se *tcp.Session, reportID string) // 友谊赛回放
	GW2CS_GuildMemberList(cmd *rpc.SimpleCmd, se *tcp.Session) // 成员列表
	GW2CS_GuildTopList(cmd *rpc.SimpleCmd, se *tcp.Session) // 公会排行榜
	GW2CS_GuildMemberDetail(cmd *rpc.SimpleCmd, se *tcp.Session, uid int64) // 查询公会成员详情
	GW2CS_WatchAds(cmd *rpc.SimpleCmd, se *tcp.Session) // 观看广告
	GW2CS_BattleReportReplay(cmd *rpc.SimpleCmd, se *tcp.Session, reportID string) // 战报回放
	GW2CS_GetUserEmails(cmd *rpc.SimpleCmd, se *tcp.Session) // 获取所有邮件
	GW2CS_GetEmailAttachment(cmd *rpc.SimpleCmd, se *tcp.Session, emailId int32) // 获取邮件的附件奖励
	GW2CS_SetEmailReaded(cmd *rpc.SimpleCmd, se *tcp.Session, emailId int32) // 设置邮件的为已读
	GW2CS_DelEmail(cmd *rpc.SimpleCmd, se *tcp.Session, emailId int32) // 删除邮件
	GW2CS_GetAllEmailAttachment(cmd *rpc.SimpleCmd, se *tcp.Session) // 获取所有邮件的附件奖励
	GW2CS_SetAllEmailAttachmentReaded(cmd *rpc.SimpleCmd, se *tcp.Session) // 设置所有邮件的为已读
	GW2CS_DelAllEmailReadedAndAttachmented(cmd *rpc.SimpleCmd, se *tcp.Session) // 删除所有已读邮件和已经领取的邮件
	GW2CS_GetAchievement(cmd *rpc.SimpleCmd, se *tcp.Session, aid int32, atype int32) // 领取成就
	GW2CS_AddComment(cmd *rpc.SimpleCmd, se *tcp.Session, hid int32, content string) // 添加评论
	GW2CS_GetComments(cmd *rpc.SimpleCmd, se *tcp.Session, hid int32, page int32, size int32) // 获取评论
	GB2CS_ReceiveComments(cmd *rpc.SimpleCmd, se *tcp.Session, wraps *wraps.HeroCommentListDataWrap) // 接收评论
	GW2CS_AddCommentLike(cmd *rpc.SimpleCmd, se *tcp.Session, hid int32, cid int32) // 对评论点赞
	GW2CS_SetTutorialMask(cmd *rpc.SimpleCmd, se *tcp.Session, data int64) // 设置新手引导
	GW2CS_GetLastBattleResult(cmd *rpc.SimpleCmd, se *tcp.Session) // 获取最后一次战斗结果
	GW2CS_GMCommand(cmd *rpc.SimpleCmd, se *tcp.Session, kind string, dataid string, count int32) // gm调试
	
}

type ServerCSCmdsInvoker struct {
	invoker IServerCSCmds
	defaultInvoker func(cmd *rpc.SimpleCmd, se *tcp.Session)
	rpc.SimpleInvoker
} 

func NewServerCSCmdsInvoker(invoker IServerCSCmds, defaultInvoker func(*rpc.SimpleCmd, *tcp.Session)) *ServerCSCmdsInvoker {
	inv := new(ServerCSCmdsInvoker)
	inv.invoker = invoker
	inv.defaultInvoker = defaultInvoker
	return inv
} 

func (this *ServerCSCmdsInvoker) Invoke(cmd *rpc.SimpleCmd, se *tcp.Session) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	pack := cmd.Pack()
	switch(cmd.Opcode()) {
	case 22001: 
		this.invoker.GW2CS_Ping(cmd,se)
	case 22002: 
		this.invoker.GS2CS_Ping(cmd,se)
	case 22003: 
		this.invoker.GB2CS_Pong(cmd,se)
	case 22004: 
		this.invoker.GW2CS_LoginGuest(cmd,se, pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString())
	case 22005: 
		this.invoker.GW2CS_LoginToken(cmd,se, pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopBool(), pack.PopString())
	case 22006: 
		this.invoker.GW2CS_BindAccount(cmd,se, pack.PopInt64(), pack.PopString(), pack.PopString())
	case 22007: 
		this.invoker.GS2CS_Register(cmd,se, pack.PopInt32(), pack.PopString(), pack.PopInt32())
	case 22008: 
		this.invoker.GW2CS_Logout(cmd,se, pack.PopInt64())
	case 22009: 
		this.invoker.GW2CS_SessionDisconnect(cmd,se, pack.PopInt64())
	case 22010: 
		this.invoker.GB2CS_RechargeNotify(cmd,se, pack.PopInt32(), pack.PopInt64(), pack.PopString(), pack.PopString(), pack.PopFloat32())
	case 22011: 
		this.invoker.GB2CS_BanUser(cmd,se, pack.PopInt32(), pack.PopInt64(), pack.PopInt32(), pack.PopInt64())
	case 22012: 
		this.invoker.GB2CS_UnBanUser(cmd,se, pack.PopInt32(), pack.PopInt64())
	case 22013: 
		this.invoker.GB2CS_CreateUserAck(cmd,se, pack.PopInt32(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString())
	case 22014: 
		this.invoker.GB2CS_UpdateUserScore(cmd,se, pack.PopInt64(), pack.PopBool(), pack.PopInt32())
	case 22015: 
		this.invoker.GB2CS_DeleteUserData(cmd,se, pack.PopInt64(), pack.PopBool())
	case 22016: 
		this.invoker.GB2CS_PvpResult(cmd,se, new(wraps.BattleReportWrap).Decode(pack).(*wraps.BattleReportWrap))
	case 22017: 
		this.invoker.GB2CS_BattleCombo(cmd,se, pack.PopInt64(), new(wraps.IntArrayWrap).Decode(pack).(*wraps.IntArrayWrap))
	case 22018: 
		this.invoker.GB2CS_SeasonTime(cmd,se, pack.PopInt32(), pack.PopInt64())
	case 22019: 
		this.invoker.GS2CS_CheckQuestResult(cmd,se, pack.PopBool(), new(wraps.BattleReportWrap).Decode(pack).(*wraps.BattleReportWrap), pack.PopInt32())
	case 22020: 
		this.invoker.GB2CS_SendSystemEmail(cmd,se, pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopInt32())
	case 22021: 
		this.invoker.GB2CS_SendPersonEmail(cmd,se, pack.PopInt64(), pack.PopString(), pack.PopString(), pack.PopString(), pack.PopString(), new(wraps.StringKVArrWrap).Decode(pack).(*wraps.StringKVArrWrap))
	case 22022: 
		this.invoker.GB2CS_GuildCreateSuccess(cmd,se)
	case 22023: 
		this.invoker.GB2CS_GuildDemandSuccess(cmd,se, pack.PopInt32())
	case 22024: 
		this.invoker.GB2CS_GuildDonateSuccess(cmd,se, pack.PopInt32(), pack.PopInt32())
	case 22025: 
		this.invoker.GB2CS_SyncTopList(cmd,se, new(wraps.TopListWrap).Decode(pack).(*wraps.TopListWrap))
	case 22026: 
		this.invoker.GB2CS_SyncMemberList(cmd,se, new(wraps.MemberListSyncWrap).Decode(pack).(*wraps.MemberListSyncWrap))
	case 22027: 
		this.invoker.GB2CS_SyncMember(cmd,se, new(wraps.MemberSyncWrap).Decode(pack).(*wraps.MemberSyncWrap))
	case 22028: 
		this.invoker.GB2CS_ReceiveDonateHero(cmd,se, pack.PopInt64(), pack.PopInt32(), pack.PopInt32())
	case 22029: 
		this.invoker.GB2CS_SetDungeon(cmd,se, new(wraps.ActivityDungeonListWrap).Decode(pack).(*wraps.ActivityDungeonListWrap))
	case 22030: 
		this.invoker.Agent2CS_Pong(cmd,se)
	case 22031: 
		this.invoker.GB2CS_SyncGuild(cmd,se, new(wraps.GuildSyncWrap).Decode(pack).(*wraps.GuildSyncWrap))
	case 12100: 
		this.invoker.GW2CS_UpBattleHeros(cmd,se, pack.PopInt32(), new(wraps.StringKVArrWrap).Decode(pack).(*wraps.StringKVArrWrap))
	case 12101: 
		this.invoker.GW2CS_ChangeBattleIndex(cmd,se, pack.PopInt32())
	case 12102: 
		this.invoker.GW2CS_HeroLevelup(cmd,se, pack.PopInt32())
	case 12103: 
		this.invoker.GW2CS_SetName(cmd,se, pack.PopString())
	case 12104: 
		this.invoker.GW2CS_SetIcon(cmd,se, pack.PopInt32())
	case 12110: 
		this.invoker.GW2CS_PvpMatch(cmd,se, pack.PopInt16(), pack.PopInt16())
	case 12111: 
		this.invoker.GW2CS_PvpCancel(cmd,se)
	case 12112: 
		this.invoker.GW2CS_StartQuest(cmd,se, pack.PopInt32(), pack.PopInt32(), pack.PopInt16())
	case 12115: 
		this.invoker.GW2CS_BuyArenaTicket(cmd,se)
	case 12116: 
		this.invoker.GW2CS_ArenaHeroOnBattle(cmd,se, pack.PopInt32(), pack.PopInt32(), pack.PopInt32(), pack.PopInt32())
	case 12117: 
		this.invoker.GW2CS_FinishQuest(cmd,se, new(wraps.BattleReportWrap).Decode(pack).(*wraps.BattleReportWrap))
	case 12119: 
		this.invoker.GW2CS_GetArenaReward(cmd,se)
	case 12120: 
		this.invoker.GW2CS_GetDailyGift(cmd,se, pack.PopInt32())
	case 12121: 
		this.invoker.GW2CS_RefreshDailyGift(cmd,se)
	case 12122: 
		this.invoker.GW2CS_GetDailyActivity(cmd,se)
	case 12130: 
		this.invoker.GW2CS_ActiveChest(cmd,se, pack.PopInt32())
	case 12131: 
		this.invoker.GW2CS_OpenChest(cmd,se, pack.PopInt32())
	case 12132: 
		this.invoker.GW2CS_OpenChestByDiamond(cmd,se, pack.PopInt32())
	case 12133: 
		this.invoker.GW2CS_OpenFreeChest(cmd,se)
	case 12134: 
		this.invoker.GW2CS_OpenArenaChest(cmd,se)
	case 12140: 
		this.invoker.GW2CS_LadderTopList(cmd,se, pack.PopInt16())
	case 12141: 
		this.invoker.GW2CS_OpenDungeon(cmd,se, pack.PopInt32())
	case 12142: 
		this.invoker.GW2CS_RestoreDungeon(cmd,se, pack.PopInt32(), pack.PopBool())
	case 12143: 
		this.invoker.GW2CS_DungeonLoot(cmd,se, pack.PopInt32())
	case 12144: 
		this.invoker.GW2CS_DungeonBattleFail(cmd,se, pack.PopInt32())
	case 12145: 
		this.invoker.GW2CS_DungeonAddHp(cmd,se, pack.PopInt32(), pack.PopInt16())
	case 12150: 
		this.invoker.GW2CS_ShopBuyChest(cmd,se, pack.PopInt32())
	case 12151: 
		this.invoker.GW2CS_ShopBuyHero(cmd,se, pack.PopInt32())
	case 12152: 
		this.invoker.GW2CS_ShopBuyGold(cmd,se, pack.PopInt32())
	case 12153: 
		this.invoker.GW2CS_ShopReloadItems(cmd,se)
	case 12154: 
		this.invoker.GW2CS_ShopBuyDiamond(cmd,se, pack.PopInt32(), pack.PopString())
	case 12160: 
		this.invoker.GW2CS_CreateGuild(cmd,se, pack.PopString(), pack.PopString(), pack.PopInt32(), pack.PopInt16(), pack.PopInt32())
	case 12161: 
		this.invoker.GW2CS_QueryGuild(cmd,se, pack.PopString(), pack.PopInt16(), pack.PopInt16(), pack.PopInt32(), pack.PopBool())
	case 12162: 
		this.invoker.GW2CS_QueryGuildDetail(cmd,se, pack.PopInt32())
	case 12163: 
		this.invoker.GW2CS_RequestJoinGuild(cmd,se, pack.PopInt32(), pack.PopString())
	case 12164: 
		this.invoker.GW2CS_DealJoinRequest(cmd,se, pack.PopInt64(), pack.PopBool())
	case 12165: 
		this.invoker.GW2CS_LeaveGuild(cmd,se)
	case 12166: 
		this.invoker.GW2CS_KickMember(cmd,se, pack.PopInt64())
	case 12167: 
		this.invoker.GW2CS_AppointMember(cmd,se, pack.PopInt64(), pack.PopInt16())
	case 12168: 
		this.invoker.GW2CS_GuildChat(cmd,se, pack.PopString())
	case 12169: 
		this.invoker.GW2CS_GuildDemand(cmd,se, pack.PopInt32())
	case 12170: 
		this.invoker.GW2CS_GuildDonate(cmd,se, pack.PopInt64(), pack.PopInt32(), pack.PopInt32())
	case 12171: 
		this.invoker.GW2CS_GuildSetting(cmd,se, pack.PopString(), pack.PopInt32(), pack.PopInt16(), pack.PopInt32())
	case 12172: 
		this.invoker.GW2CS_GuildRequestPvP(cmd,se, pack.PopString(), pack.PopInt16())
	case 12173: 
		this.invoker.GW2CS_GuildCancelPvP(cmd,se)
	case 12174: 
		this.invoker.GW2CS_GuildStartPvP(cmd,se, pack.PopInt64(), pack.PopInt16())
	case 12175: 
		this.invoker.GW2CS_GuildPvPReplay(cmd,se, pack.PopString())
	case 12176: 
		this.invoker.GW2CS_GuildMemberList(cmd,se)
	case 12177: 
		this.invoker.GW2CS_GuildTopList(cmd,se)
	case 12178: 
		this.invoker.GW2CS_GuildMemberDetail(cmd,se, pack.PopInt64())
	case 12180: 
		this.invoker.GW2CS_WatchAds(cmd,se)
	case 12182: 
		this.invoker.GW2CS_BattleReportReplay(cmd,se, pack.PopString())
	case 12192: 
		this.invoker.GW2CS_GetUserEmails(cmd,se)
	case 12193: 
		this.invoker.GW2CS_GetEmailAttachment(cmd,se, pack.PopInt32())
	case 12194: 
		this.invoker.GW2CS_SetEmailReaded(cmd,se, pack.PopInt32())
	case 12195: 
		this.invoker.GW2CS_DelEmail(cmd,se, pack.PopInt32())
	case 12196: 
		this.invoker.GW2CS_GetAllEmailAttachment(cmd,se)
	case 12197: 
		this.invoker.GW2CS_SetAllEmailAttachmentReaded(cmd,se)
	case 12198: 
		this.invoker.GW2CS_DelAllEmailReadedAndAttachmented(cmd,se)
	case 12200: 
		this.invoker.GW2CS_GetAchievement(cmd,se, pack.PopInt32(), pack.PopInt32())
	case 12301: 
		this.invoker.GW2CS_AddComment(cmd,se, pack.PopInt32(), pack.PopString())
	case 12302: 
		this.invoker.GW2CS_GetComments(cmd,se, pack.PopInt32(), pack.PopInt32(), pack.PopInt32())
	case 12303: 
		this.invoker.GB2CS_ReceiveComments(cmd,se, new(wraps.HeroCommentListDataWrap).Decode(pack).(*wraps.HeroCommentListDataWrap))
	case 12304: 
		this.invoker.GW2CS_AddCommentLike(cmd,se, pack.PopInt32(), pack.PopInt32())
	case 12310: 
		this.invoker.GW2CS_SetTutorialMask(cmd,se, pack.PopInt64())
	case 12311: 
		this.invoker.GW2CS_GetLastBattleResult(cmd,se)
	case 12999: 
		this.invoker.GW2CS_GMCommand(cmd,se, pack.PopString(), pack.PopString(), pack.PopInt32())
	
	default:
		if this.defaultInvoker != nil {
			this.defaultInvoker(cmd,se)
		}
	}
	return nil
}

