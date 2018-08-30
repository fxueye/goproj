package wraps

import (
	rpc "game/common/rpc/simple"
)

type PlayerWrap struct {
	
	GUID string // 唯一ID
	CreateTime int64 // 创建时间
	rpc.Wrapper
}

func (w *PlayerWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GUID = pck.PopString(); 
	w.CreateTime = pck.PopInt64(); 
	return w
}

func (w *PlayerWrap)Encode(pck *rpc.Packet) {
	
	pck.PutString(w.GUID) 
	pck.PutInt64(w.CreateTime) 
}

