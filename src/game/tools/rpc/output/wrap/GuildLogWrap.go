package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildLogWrap struct {
	
	Opcode int16 // 操作类型
	CreateTime int64 // 时间
	MemberID int64 // 操作者ID
	MemberName string // 操作者名字
	TargetID int64 // 目标ID
	TargetName string // 目标名字
	GuildRank int16 // 操作官职
	rpc.Wrapper
}

func (w *GuildLogWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Opcode = pck.PopInt16(); 
	w.CreateTime = pck.PopInt64(); 
	w.MemberID = pck.PopInt64(); 
	w.MemberName = pck.PopString(); 
	w.TargetID = pck.PopInt64(); 
	w.TargetName = pck.PopString(); 
	w.GuildRank = pck.PopInt16(); 
	return w
}

func (w *GuildLogWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt16(w.Opcode) 
	pck.PutInt64(w.CreateTime) 
	pck.PutInt64(w.MemberID) 
	pck.PutString(w.MemberName) 
	pck.PutInt64(w.TargetID) 
	pck.PutString(w.TargetName) 
	pck.PutInt16(w.GuildRank) 
}

