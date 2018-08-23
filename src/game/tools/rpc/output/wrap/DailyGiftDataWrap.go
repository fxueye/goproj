package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DailyGiftDataWrap struct {
	
	GiftIdx int32 // 数据下标
	GiftID int32 // 礼品ID
	GiftType int32 // 类型 1是item 2是hero
	Count int32 // 数量
	State int32 // 状态 0已经领取 1可以领取 2不可以领取 3 补签  4 已经补签
	GroupID int32 // 组ID
	rpc.Wrapper
}

func (w *DailyGiftDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GiftIdx = pck.PopInt32(); 
	w.GiftID = pck.PopInt32(); 
	w.GiftType = pck.PopInt32(); 
	w.Count = pck.PopInt32(); 
	w.State = pck.PopInt32(); 
	w.GroupID = pck.PopInt32(); 
	return w
}

func (w *DailyGiftDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GiftIdx) 
	pck.PutInt32(w.GiftID) 
	pck.PutInt32(w.GiftType) 
	pck.PutInt32(w.Count) 
	pck.PutInt32(w.State) 
	pck.PutInt32(w.GroupID) 
}

