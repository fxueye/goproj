package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type BattleReportListWrap struct {
	
	List []*BattleReportWrap // 列表
	rpc.Wrapper
}

func (w *BattleReportListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.List = make([]*BattleReportWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(BattleReportWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *BattleReportListWrap)Encode(pck *rpc.Packet) {
	
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

