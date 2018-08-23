package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildMemberListWrap struct {
	
	GuildID int32 // 公会id
	Members []*GuildMemberBriefWrap // 成员
	rpc.Wrapper
}

func (w *GuildMemberListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GuildID = pck.PopInt32(); 
	w.Members = make([]*GuildMemberBriefWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Members); i++ {
		w.Members[i] = new(GuildMemberBriefWrap)
		w.Members[i].Decode(pck)
		
	}
	
	return w
}

func (w *GuildMemberListWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GuildID) 
	if w.Members == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Members)))
    	for i := 0; i < len(w.Members); i++ {
    		w.Members[i].Encode(pck);
			
    	}
    }
	
}

