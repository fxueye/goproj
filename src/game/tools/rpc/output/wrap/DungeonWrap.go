package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type DungeonWrap struct {
	
	ID int32 // 活动本id
	OpenCnt int16 // 当前开启次数
	RestoreCount int32 // 当前续命次数
	HP int32 // 当前血量
	MaxHP int32 // 最大血量
	Step int32 // 当前关卡
	DropHero []int32 // 掉落英雄
	DropCount []int32 // 掉落数量（与掉落英雄对应）
	rpc.Wrapper
}

func (w *DungeonWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ID = pck.PopInt32(); 
	w.OpenCnt = pck.PopInt16(); 
	w.RestoreCount = pck.PopInt32(); 
	w.HP = pck.PopInt32(); 
	w.MaxHP = pck.PopInt32(); 
	w.Step = pck.PopInt32(); 
	w.DropHero = make([]int32, int(pck.PopInt16()))
	for i := 0; i < len(w.DropHero); i++ {
		w.DropHero[i] = pck.PopInt32()
	}
	
	w.DropCount = make([]int32, int(pck.PopInt16()))
	for i := 0; i < len(w.DropCount); i++ {
		w.DropCount[i] = pck.PopInt32()
	}
	
	return w
}

func (w *DungeonWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ID) 
	pck.PutInt16(w.OpenCnt) 
	pck.PutInt32(w.RestoreCount) 
	pck.PutInt32(w.HP) 
	pck.PutInt32(w.MaxHP) 
	pck.PutInt32(w.Step) 
	if w.DropHero == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.DropHero)))
    	for i := 0; i < len(w.DropHero); i++ {
    		pck.PutInt32(w.DropHero[i])
    	}
    }
	
	if w.DropCount == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.DropCount)))
    	for i := 0; i < len(w.DropCount); i++ {
    		pck.PutInt32(w.DropCount[i])
    	}
    }
	
}

