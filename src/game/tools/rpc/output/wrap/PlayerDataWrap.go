package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type PlayerDataWrap struct {
	
	DataKey []string // 数据Key
	DataValue []int64 // 数据Value
	rpc.Wrapper
}

func (w *PlayerDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.DataKey = make([]string, int(pck.PopInt16()))
	for i := 0; i < len(w.DataKey); i++ {
		w.DataKey[i] = pck.PopString()
	}
	
	w.DataValue = make([]int64, int(pck.PopInt16()))
	for i := 0; i < len(w.DataValue); i++ {
		w.DataValue[i] = pck.PopInt64()
	}
	
	return w
}

func (w *PlayerDataWrap)Encode(pck *rpc.Packet) {
	
	if w.DataKey == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.DataKey)))
    	for i := 0; i < len(w.DataKey); i++ {
    		pck.PutString(w.DataKey[i])
    	}
    }
	
	if w.DataValue == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.DataValue)))
    	for i := 0; i < len(w.DataValue); i++ {
    		pck.PutInt64(w.DataValue[i])
    	}
    }
	
}

