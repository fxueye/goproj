package wraps

import (
	rpc "game/common/rpc/simple"
)

type ItemDataWrap struct {
	
	ItemID int32 // 道具ID
	Count int32 // 道具数量
	LastUpdateTime int64 // 最后一次更新时间
	rpc.Wrapper
}

func (w *ItemDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ItemID = pck.PopInt32(); 
	w.Count = pck.PopInt32(); 
	w.LastUpdateTime = pck.PopInt64(); 
	return w
}

func (w *ItemDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ItemID) 
	pck.PutInt32(w.Count) 
	pck.PutInt64(w.LastUpdateTime) 
}

