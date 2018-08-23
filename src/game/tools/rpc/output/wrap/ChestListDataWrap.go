package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ChestListDataWrap struct {
	
	ChestList []*ChestDataWrap // 宝箱列表
	rpc.Wrapper
}

func (w *ChestListDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ChestList = make([]*ChestDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.ChestList); i++ {
		w.ChestList[i] = new(ChestDataWrap)
		w.ChestList[i].Decode(pck)
		
	}
	
	return w
}

func (w *ChestListDataWrap)Encode(pck *rpc.Packet) {
	
	if w.ChestList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.ChestList)))
    	for i := 0; i < len(w.ChestList); i++ {
    		w.ChestList[i].Encode(pck);
			
    	}
    }
	
}

