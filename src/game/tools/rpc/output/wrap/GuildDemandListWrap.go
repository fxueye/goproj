package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildDemandListWrap struct {
	
	List []*GuildDemandWrap // 列表
	rpc.Wrapper
}

func (w *GuildDemandListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.List = make([]*GuildDemandWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(GuildDemandWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *GuildDemandListWrap)Encode(pck *rpc.Packet) {
	
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

