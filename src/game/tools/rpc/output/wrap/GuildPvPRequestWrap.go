package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildPvPRequestWrap struct {
	
	MemberID int64 // 会员id
	Name string // 名字
	IconID int32 // 头像ID
	Msg string // 留言内容
	rpc.Wrapper
}

func (w *GuildPvPRequestWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.MemberID = pck.PopInt64(); 
	w.Name = pck.PopString(); 
	w.IconID = pck.PopInt32(); 
	w.Msg = pck.PopString(); 
	return w
}

func (w *GuildPvPRequestWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.MemberID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.IconID) 
	pck.PutString(w.Msg) 
}

