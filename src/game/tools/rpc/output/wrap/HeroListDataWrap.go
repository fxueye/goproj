package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type HeroListDataWrap struct {
	
	HerosList []*HeroDataWrap // 英雄列表
	rpc.Wrapper
}

func (w *HeroListDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.HerosList = make([]*HeroDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.HerosList); i++ {
		w.HerosList[i] = new(HeroDataWrap)
		w.HerosList[i].Decode(pck)
		
	}
	
	return w
}

func (w *HeroListDataWrap)Encode(pck *rpc.Packet) {
	
	if w.HerosList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.HerosList)))
    	for i := 0; i < len(w.HerosList); i++ {
    		w.HerosList[i].Encode(pck);
			
    	}
    }
	
}

