package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type BattlePlayerWrap struct {
	
	UID int64 // 游戏的uid
	UserName string // 用户名（唯一，初始和UID一致）
	Score int32 // 积分
	Level int32 // 等级
	Icon int32 // 头像
	StartHP int32 // 初始血量
	Heroes []*BattleHeroWrap // 上阵英雄
	rpc.Wrapper
}

func (w *BattlePlayerWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.UserName = pck.PopString(); 
	w.Score = pck.PopInt32(); 
	w.Level = pck.PopInt32(); 
	w.Icon = pck.PopInt32(); 
	w.StartHP = pck.PopInt32(); 
	w.Heroes = make([]*BattleHeroWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Heroes); i++ {
		w.Heroes[i] = new(BattleHeroWrap)
		w.Heroes[i].Decode(pck)
		
	}
	
	return w
}

func (w *BattlePlayerWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutString(w.UserName) 
	pck.PutInt32(w.Score) 
	pck.PutInt32(w.Level) 
	pck.PutInt32(w.Icon) 
	pck.PutInt32(w.StartHP) 
	if w.Heroes == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Heroes)))
    	for i := 0; i < len(w.Heroes); i++ {
    		w.Heroes[i].Encode(pck);
			
    	}
    }
	
}

