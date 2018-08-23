package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type BattleCmdWrap struct {
	
	SeqID int32 // 序列号
	Opcode int16 // 协议号
	Timestamp int64 // 时间戳
	Args []int32 // 参数数组
	rpc.Wrapper
}

func (w *BattleCmdWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.SeqID = pck.PopInt32(); 
	w.Opcode = pck.PopInt16(); 
	w.Timestamp = pck.PopInt64(); 
	w.Args = make([]int32, int(pck.PopInt16()))
	for i := 0; i < len(w.Args); i++ {
		w.Args[i] = pck.PopInt32()
	}
	
	return w
}

func (w *BattleCmdWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.SeqID) 
	pck.PutInt16(w.Opcode) 
	pck.PutInt64(w.Timestamp) 
	if w.Args == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Args)))
    	for i := 0; i < len(w.Args); i++ {
    		pck.PutInt32(w.Args[i])
    	}
    }
	
}

