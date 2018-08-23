package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type PlayerLevelUpWrap struct {
	
	oldLv int32 // 旧等级
	newLv int32 // 新等级
	oldBlood int32 // 旧血量
	newBlood int32 // 新血量
	diamondAdd int32 // 获得钻石
	coinAdd int32 // 获得金币
	rpc.Wrapper
}

func (w *PlayerLevelUpWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.oldLv = pck.PopInt32(); 
	w.newLv = pck.PopInt32(); 
	w.oldBlood = pck.PopInt32(); 
	w.newBlood = pck.PopInt32(); 
	w.diamondAdd = pck.PopInt32(); 
	w.coinAdd = pck.PopInt32(); 
	return w
}

func (w *PlayerLevelUpWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.oldLv) 
	pck.PutInt32(w.newLv) 
	pck.PutInt32(w.oldBlood) 
	pck.PutInt32(w.newBlood) 
	pck.PutInt32(w.diamondAdd) 
	pck.PutInt32(w.coinAdd) 
}

