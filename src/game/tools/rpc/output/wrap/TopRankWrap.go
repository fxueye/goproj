package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type TopRankWrap struct {
	
	UID int64 // UID
	Name string // 名字
	IconID int32 // 头像ID
	Data []int32 // 数据
	rpc.Wrapper
}

func (w *TopRankWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.Name = pck.PopString(); 
	w.IconID = pck.PopInt32(); 
	w.Data = make([]int32, int(pck.PopInt16()))
	for i := 0; i < len(w.Data); i++ {
		w.Data[i] = pck.PopInt32()
	}
	
	return w
}

func (w *TopRankWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.IconID) 
	if w.Data == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Data)))
    	for i := 0; i < len(w.Data); i++ {
    		pck.PutInt32(w.Data[i])
    	}
    }
	
}

