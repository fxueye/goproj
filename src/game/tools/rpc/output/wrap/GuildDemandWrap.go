package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildDemandWrap struct {
	
	MemberID int64 // 会员id
	Name string // 名字
	IconID int32 // 头像ID
	DemandTime int64 // 请求时间
	ExpireTime int64 // 过期时间
	HeroID int32 // 索要英雄
	MaxCount int32 // 最大数量
	CurCount int32 // 当前数量
	rpc.Wrapper
}

func (w *GuildDemandWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.MemberID = pck.PopInt64(); 
	w.Name = pck.PopString(); 
	w.IconID = pck.PopInt32(); 
	w.DemandTime = pck.PopInt64(); 
	w.ExpireTime = pck.PopInt64(); 
	w.HeroID = pck.PopInt32(); 
	w.MaxCount = pck.PopInt32(); 
	w.CurCount = pck.PopInt32(); 
	return w
}

func (w *GuildDemandWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.MemberID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.IconID) 
	pck.PutInt64(w.DemandTime) 
	pck.PutInt64(w.ExpireTime) 
	pck.PutInt32(w.HeroID) 
	pck.PutInt32(w.MaxCount) 
	pck.PutInt32(w.CurCount) 
}

