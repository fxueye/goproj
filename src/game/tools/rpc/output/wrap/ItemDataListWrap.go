package wraps

import (
	rpc "game/common/rpc/simple"
)

type ItemDataListWrap struct {
	
	ItemList []*ItemDataWrap // 道具列表
	rpc.Wrapper
}

func (w *ItemDataListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ItemList = make([]*ItemDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.ItemList); i++ {
		w.ItemList[i] = new(ItemDataWrap)
		w.ItemList[i].Decode(pck)
		
	}
	
	return w
}

func (w *ItemDataListWrap)Encode(pck *rpc.Packet) {
	
	if w.ItemList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.ItemList)))
    	for i := 0; i < len(w.ItemList); i++ {
    		w.ItemList[i].Encode(pck);
			
    	}
    }
	
}

