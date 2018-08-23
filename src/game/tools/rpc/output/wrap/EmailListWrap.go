package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type EmailListWrap struct {
	
	List []*EmailWrap // 数据
	rpc.Wrapper
}

func (w *EmailListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.List = make([]*EmailWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.List); i++ {
		w.List[i] = new(EmailWrap)
		w.List[i].Decode(pck)
		
	}
	
	return w
}

func (w *EmailListWrap)Encode(pck *rpc.Packet) {
	
	if w.List == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.List)))
    	for i := 0; i < len(w.List); i++ {
    		w.List[i].Encode(pck);
			
    	}
    }
	
}

