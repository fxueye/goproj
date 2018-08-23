package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type StringKVArrWrap struct {
	
	StringMap []*StringKVWrap // map<string,string>
	rpc.Wrapper
}

func (w *StringKVArrWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.StringMap = make([]*StringKVWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.StringMap); i++ {
		w.StringMap[i] = new(StringKVWrap)
		w.StringMap[i].Decode(pck)
		
	}
	
	return w
}

func (w *StringKVArrWrap)Encode(pck *rpc.Packet) {
	
	if w.StringMap == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.StringMap)))
    	for i := 0; i < len(w.StringMap); i++ {
    		w.StringMap[i].Encode(pck);
			
    	}
    }
	
}

