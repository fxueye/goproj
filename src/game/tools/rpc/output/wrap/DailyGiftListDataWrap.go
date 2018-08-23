package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DailyGiftListDataWrap struct {
	
	DailyList []*DailyGiftDataWrap // 日常数据
	NextTime int64 // 下次领取新时间
	ActivityList []*DailyActivityDataWrap // 活跃度数据
	rpc.Wrapper
}

func (w *DailyGiftListDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.DailyList = make([]*DailyGiftDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.DailyList); i++ {
		w.DailyList[i] = new(DailyGiftDataWrap)
		w.DailyList[i].Decode(pck)
		
	}
	
	w.NextTime = pck.PopInt64(); 
	w.ActivityList = make([]*DailyActivityDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.ActivityList); i++ {
		w.ActivityList[i] = new(DailyActivityDataWrap)
		w.ActivityList[i].Decode(pck)
		
	}
	
	return w
}

func (w *DailyGiftListDataWrap)Encode(pck *rpc.Packet) {
	
	if w.DailyList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.DailyList)))
    	for i := 0; i < len(w.DailyList); i++ {
    		w.DailyList[i].Encode(pck);
			
    	}
    }
	
	pck.PutInt64(w.NextTime) 
	if w.ActivityList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.ActivityList)))
    	for i := 0; i < len(w.ActivityList); i++ {
    		w.ActivityList[i].Encode(pck);
			
    	}
    }
	
}

