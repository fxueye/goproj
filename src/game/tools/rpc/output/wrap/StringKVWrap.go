package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type StringKVWrap struct {
	
	Key string // map的key
	Value string // map的value
	rpc.Wrapper
}

func (w *StringKVWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Key = pck.PopString(); 
	w.Value = pck.PopString(); 
	return w
}

func (w *StringKVWrap)Encode(pck *rpc.Packet) {
	
	pck.PutString(w.Key) 
	pck.PutString(w.Value) 
}

