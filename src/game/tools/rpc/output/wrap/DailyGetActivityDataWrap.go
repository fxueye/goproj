package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DailyGetActivityDataWrap struct {
	
	ActIdxs []int32 // 当前领取的id数组
	ChestDataWrap *ChestOpenDataWrap // 领取的结果数据
	rpc.Wrapper
}

func (w *DailyGetActivityDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ActIdxs = make([]int32, int(pck.PopInt16()))
	for i := 0; i < len(w.ActIdxs); i++ {
		w.ActIdxs[i] = pck.PopInt32()
	}
	
	w.ChestDataWrap = new(ChestOpenDataWrap)
	w.ChestDataWrap.Decode(pck)
	
	return w
}

func (w *DailyGetActivityDataWrap)Encode(pck *rpc.Packet) {
	
	if w.ActIdxs == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.ActIdxs)))
    	for i := 0; i < len(w.ActIdxs); i++ {
    		pck.PutInt32(w.ActIdxs[i])
    	}
    }
	
	w.ChestDataWrap.Encode(pck); 
}

