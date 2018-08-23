package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildBriefWrap struct {
	
	GuildID int32 // 公会id
	Name string // 名字
	IconID int32 // 头像ID
	Member int16 // 成员数
	Score int32 // 公会积分
	rpc.Wrapper
}

func (w *GuildBriefWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GuildID = pck.PopInt32(); 
	w.Name = pck.PopString(); 
	w.IconID = pck.PopInt32(); 
	w.Member = pck.PopInt16(); 
	w.Score = pck.PopInt32(); 
	return w
}

func (w *GuildBriefWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GuildID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.IconID) 
	pck.PutInt16(w.Member) 
	pck.PutInt32(w.Score) 
}

