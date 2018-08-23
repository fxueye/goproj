package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ChestDataWrap struct {
	
	ChestIdx int32 // 宝箱位置
	ChestID int32 // 宝箱ID
	QuickenCount int32 // 加速次数
	ActiveTime int64 // 激活时间
	Lv int32 // 宝箱等级
	rpc.Wrapper
}

func (w *ChestDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ChestIdx = pck.PopInt32(); 
	w.ChestID = pck.PopInt32(); 
	w.QuickenCount = pck.PopInt32(); 
	w.ActiveTime = pck.PopInt64(); 
	w.Lv = pck.PopInt32(); 
	return w
}

func (w *ChestDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ChestIdx) 
	pck.PutInt32(w.ChestID) 
	pck.PutInt32(w.QuickenCount) 
	pck.PutInt64(w.ActiveTime) 
	pck.PutInt32(w.Lv) 
}

