package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type QuestDataListWrap struct {
	
	QuestList []*QuestDataWrap // 关卡列表
	rpc.Wrapper
}

func (w *QuestDataListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.QuestList = make([]*QuestDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.QuestList); i++ {
		w.QuestList[i] = new(QuestDataWrap)
		w.QuestList[i].Decode(pck)
		
	}
	
	return w
}

func (w *QuestDataListWrap)Encode(pck *rpc.Packet) {
	
	if w.QuestList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.QuestList)))
    	for i := 0; i < len(w.QuestList); i++ {
    		w.QuestList[i].Encode(pck);
			
    	}
    }
	
}

