package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type QuestDataWrap struct {
	
	QuestID int32 // 关卡ID
	ChestState bool // 宝箱是否发放
	LastUpdateTime int64 // 最后一次更新时间
	rpc.Wrapper
}

func (w *QuestDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.QuestID = pck.PopInt32(); 
	w.ChestState = pck.PopBool(); 
	w.LastUpdateTime = pck.PopInt64(); 
	return w
}

func (w *QuestDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.QuestID) 
	pck.PutBool(w.ChestState) 
	pck.PutInt64(w.LastUpdateTime) 
}

