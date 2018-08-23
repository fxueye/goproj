package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type TopListWrap struct {
	
	Kind int16 // 类型
	List []*TopRankWrap // 数据
	rpc.Wrapper
}

func (w *TopListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Kind = pck.PopInt16(); 
	w.List = make([]*TopRankWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(TopRankWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *TopListWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt16(w.Kind) 
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

