package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type IntKVWrap struct {
	
	Key int16 // map的key
	Value int16 // map的value
	rpc.Wrapper
}

func (w *IntKVWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Key = pck.PopInt16(); 
	w.Value = pck.PopInt16(); 
	return w
}

func (w *IntKVWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt16(w.Key) 
	pck.PutInt16(w.Value) 
}

