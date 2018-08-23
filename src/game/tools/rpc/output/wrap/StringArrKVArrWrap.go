package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type StringArrKVArrWrap struct {
	
	StringMap []*StringKVArrWrap // map<string,string>
	rpc.Wrapper
}

func (w *StringArrKVArrWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.StringMap = make([]*StringKVArrWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.StringMap); i++ {
		w.StringMap[i] = new(StringKVArrWrap)
		w.StringMap[i].Decode(pck)
		
	}
	
	return w
}

func (w *StringArrKVArrWrap)Encode(pck *rpc.Packet) {
	
	if w.StringMap == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.StringMap)))
    	for i := 0; i < len(w.StringMap); i++ {
    		w.StringMap[i].Encode(pck);
			
    	}
    }
	
}

