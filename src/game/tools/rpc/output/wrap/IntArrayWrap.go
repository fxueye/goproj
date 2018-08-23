package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type IntArrayWrap struct {
	
	Value []int32 // 整型数组
	rpc.Wrapper
}

func (w *IntArrayWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Value = make([]int32, int(pck.PopInt16()))
	for i := 0; i < len(w.Value); i++ {
		w.Value[i] = pck.PopInt32()
	}
	
	return w
}

func (w *IntArrayWrap)Encode(pck *rpc.Packet) {
	
	if w.Value == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Value)))
    	for i := 0; i < len(w.Value); i++ {
    		pck.PutInt32(w.Value[i])
    	}
    }
	
}

