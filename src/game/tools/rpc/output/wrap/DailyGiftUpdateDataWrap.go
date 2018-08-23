package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DailyGiftUpdateDataWrap struct {
	
	GiftIdx int32 // 当前领取的idx
	Status int32 // 更新状态
	rpc.Wrapper
}

func (w *DailyGiftUpdateDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GiftIdx = pck.PopInt32(); 
	w.Status = pck.PopInt32(); 
	return w
}

func (w *DailyGiftUpdateDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GiftIdx) 
	pck.PutInt32(w.Status) 
}

