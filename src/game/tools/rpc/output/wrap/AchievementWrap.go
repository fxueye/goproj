package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type AchievementWrap struct {
	
	AchType int32 // 成就类型
	AchId int32 // 成就id
	AchScore int32 // 成就得分
	AchState int32 // 成就状态  0进行中 1完成 2领取
	LastTime int64 // 成就最后更新时间
	rpc.Wrapper
}

func (w *AchievementWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.AchType = pck.PopInt32(); 
	w.AchId = pck.PopInt32(); 
	w.AchScore = pck.PopInt32(); 
	w.AchState = pck.PopInt32(); 
	w.LastTime = pck.PopInt64(); 
	return w
}

func (w *AchievementWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.AchType) 
	pck.PutInt32(w.AchId) 
	pck.PutInt32(w.AchScore) 
	pck.PutInt32(w.AchState) 
	pck.PutInt64(w.LastTime) 
}

