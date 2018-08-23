package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type BattleHeroWrap struct {
	
	HeroID int32 // 英雄ID
	Level int32 // 英雄等级
	rpc.Wrapper
}

func (w *BattleHeroWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.HeroID = pck.PopInt32(); 
	w.Level = pck.PopInt32(); 
	return w
}

func (w *BattleHeroWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.HeroID) 
	pck.PutInt32(w.Level) 
}

