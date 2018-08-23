package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ChestOpenDataWrap struct {
	
	ChestIdx int32 // 宝箱位置
	ChestID int32 // 宝箱ID
	Coin int32 // 金币数量
	Diamond int32 // 钻石数量
	Heros []*HeroDataWrap // 英雄
	LadderLv int32 // 天体等级
	rpc.Wrapper
}

func (w *ChestOpenDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ChestIdx = pck.PopInt32(); 
	w.ChestID = pck.PopInt32(); 
	w.Coin = pck.PopInt32(); 
	w.Diamond = pck.PopInt32(); 
	w.Heros = make([]*HeroDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Heros); i++ {
		w.Heros[i] = new(HeroDataWrap)
		w.Heros[i].Decode(pck)
		
	}
	
	w.LadderLv = pck.PopInt32(); 
	return w
}

func (w *ChestOpenDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ChestIdx) 
	pck.PutInt32(w.ChestID) 
	pck.PutInt32(w.Coin) 
	pck.PutInt32(w.Diamond) 
	if w.Heros == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Heros)))
    	for i := 0; i < len(w.Heros); i++ {
    		w.Heros[i].Encode(pck);
			
    	}
    }
	
	pck.PutInt32(w.LadderLv) 
}

