package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ActivityDungeonWrap struct {
	
	GroupID int32 // 活动本groupid
	StartTime int64 // 开始时间
	EndTime int64 // 结束时间
	rpc.Wrapper
}

func (w *ActivityDungeonWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GroupID = pck.PopInt32(); 
	w.StartTime = pck.PopInt64(); 
	w.EndTime = pck.PopInt64(); 
	return w
}

func (w *ActivityDungeonWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GroupID) 
	pck.PutInt64(w.StartTime) 
	pck.PutInt64(w.EndTime) 
}

