package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildChatListWrap struct {
	
	List []*GuildChatWrap // 数据
	rpc.Wrapper
}

func (w *GuildChatListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.List = make([]*GuildChatWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(GuildChatWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *GuildChatListWrap)Encode(pck *rpc.Packet) {
	
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

