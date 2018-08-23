package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type PlayerDetailWrap struct {
	
	UID int64 // 游戏的uid
	UserName string // 用户名
	Icon int32 // 头像
	GuildID int32 // 公会ID
	GuildName string // 公会名
	GuildIcon int32 // 公会图标
	GuildRank int16 // 公会官阶
	Level int32 // 等级
	Exp int32 // 经验
	HeroCnt int32 // 英雄数量
	FavorHero int32 // 最常用英雄
	LadderRank int32 // 天梯名次
	MaxScore int32 // 天梯最高积分
	Score int32 // 天梯积分
	ArenaWin int32 // 竞技场最高胜场
	DonateCnt int32 // 赠送武将个数
	Heroes []*BattleHeroWrap // 上阵英雄
	rpc.Wrapper
}

func (w *PlayerDetailWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.UserName = pck.PopString(); 
	w.Icon = pck.PopInt32(); 
	w.GuildID = pck.PopInt32(); 
	w.GuildName = pck.PopString(); 
	w.GuildIcon = pck.PopInt32(); 
	w.GuildRank = pck.PopInt16(); 
	w.Level = pck.PopInt32(); 
	w.Exp = pck.PopInt32(); 
	w.HeroCnt = pck.PopInt32(); 
	w.FavorHero = pck.PopInt32(); 
	w.LadderRank = pck.PopInt32(); 
	w.MaxScore = pck.PopInt32(); 
	w.Score = pck.PopInt32(); 
	w.ArenaWin = pck.PopInt32(); 
	w.DonateCnt = pck.PopInt32(); 
	w.Heroes = make([]*BattleHeroWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Heroes); i++ {
		w.Heroes[i] = new(BattleHeroWrap)
		w.Heroes[i].Decode(pck)
		
	}
	
	return w
}

func (w *PlayerDetailWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutString(w.UserName) 
	pck.PutInt32(w.Icon) 
	pck.PutInt32(w.GuildID) 
	pck.PutString(w.GuildName) 
	pck.PutInt32(w.GuildIcon) 
	pck.PutInt16(w.GuildRank) 
	pck.PutInt32(w.Level) 
	pck.PutInt32(w.Exp) 
	pck.PutInt32(w.HeroCnt) 
	pck.PutInt32(w.FavorHero) 
	pck.PutInt32(w.LadderRank) 
	pck.PutInt32(w.MaxScore) 
	pck.PutInt32(w.Score) 
	pck.PutInt32(w.ArenaWin) 
	pck.PutInt32(w.DonateCnt) 
	if w.Heroes == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Heroes)))
    	for i := 0; i < len(w.Heroes); i++ {
    		w.Heroes[i].Encode(pck);
			
    	}
    }
	
}

