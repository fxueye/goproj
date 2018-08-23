package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type AchievementListWrap struct {
	
	AchList []*AchievementWrap // 成就列表
	rpc.Wrapper
}

func (w *AchievementListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.AchList = make([]*AchievementWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.AchList); i++ {
		w.AchList[i] = new(AchievementWrap)
		w.AchList[i].Decode(pck)
		
	}
	
	return w
}

func (w *AchievementListWrap)Encode(pck *rpc.Packet) {
	
	if w.AchList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.AchList)))
    	for i := 0; i < len(w.AchList); i++ {
    		w.AchList[i].Encode(pck);
			
    	}
    }
	
}

