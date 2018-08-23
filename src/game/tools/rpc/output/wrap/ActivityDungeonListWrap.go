package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ActivityDungeonListWrap struct {
	
	List []*ActivityDungeonWrap // 活动本列表
	rpc.Wrapper
}

func (w *ActivityDungeonListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.List = make([]*ActivityDungeonWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(ActivityDungeonWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *ActivityDungeonListWrap)Encode(pck *rpc.Packet) {
	
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

