package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DungeonListWrap struct {
	
	List []*DungeonWrap // 数据
	rpc.Wrapper
}

func (w *DungeonListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.List = make([]*DungeonWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(DungeonWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *DungeonListWrap)Encode(pck *rpc.Packet) {
	
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

