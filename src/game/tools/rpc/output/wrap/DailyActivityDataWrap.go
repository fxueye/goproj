package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DailyActivityDataWrap struct {
	
	ActivityIdx int32 // 数据下标
	ChestID int32 // 宝箱ID
	ActivityPoints int32 // 活跃度
	Lv int32 // 等级
	State int32 // 状态 0已经领取 1可以领取 2不可以领取
	rpc.Wrapper
}

func (w *DailyActivityDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ActivityIdx = pck.PopInt32(); 
	w.ChestID = pck.PopInt32(); 
	w.ActivityPoints = pck.PopInt32(); 
	w.Lv = pck.PopInt32(); 
	w.State = pck.PopInt32(); 
	return w
}

func (w *DailyActivityDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ActivityIdx) 
	pck.PutInt32(w.ChestID) 
	pck.PutInt32(w.ActivityPoints) 
	pck.PutInt32(w.Lv) 
	pck.PutInt32(w.State) 
}

