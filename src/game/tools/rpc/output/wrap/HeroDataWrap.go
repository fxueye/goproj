package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type HeroDataWrap struct {
	
	HeroID int32 // 卡牌ID
	Count int32 // 卡牌数量
	Level int32 // 等级
	LastUpdateTime int64 // 最后一次更新时间
	UseCnt int32 // 使用次数
	rpc.Wrapper
}

func (w *HeroDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.HeroID = pck.PopInt32(); 
	w.Count = pck.PopInt32(); 
	w.Level = pck.PopInt32(); 
	w.LastUpdateTime = pck.PopInt64(); 
	w.UseCnt = pck.PopInt32(); 
	return w
}

func (w *HeroDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.HeroID) 
	pck.PutInt32(w.Count) 
	pck.PutInt32(w.Level) 
	pck.PutInt64(w.LastUpdateTime) 
	pck.PutInt32(w.UseCnt) 
}

