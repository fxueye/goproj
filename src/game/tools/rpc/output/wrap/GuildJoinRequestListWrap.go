package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildJoinRequestListWrap struct {
	
	Requests []*GuildJoinRequestWrap // 申请列表
	rpc.Wrapper
}

func (w *GuildJoinRequestListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Requests = make([]*GuildJoinRequestWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Requests); i++ {
		w.Requests[i] = new(GuildJoinRequestWrap)
		w.Requests[i].Decode(pck)
		
	}
	
	return w
}

func (w *GuildJoinRequestListWrap)Encode(pck *rpc.Packet) {
	
	if w.Requests == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Requests)))
    	for i := 0; i < len(w.Requests); i++ {
    		w.Requests[i].Encode(pck);
			
    	}
    }
	
}

