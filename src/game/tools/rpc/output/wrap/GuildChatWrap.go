package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildChatWrap struct {
	
	MemberID int64 // 会员id
	Name string // 名字
	IconID int32 // 头像ID
	ChatTime int64 // 聊天时间
	Msg string // 聊天内容
	rpc.Wrapper
}

func (w *GuildChatWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.MemberID = pck.PopInt64(); 
	w.Name = pck.PopString(); 
	w.IconID = pck.PopInt32(); 
	w.ChatTime = pck.PopInt64(); 
	w.Msg = pck.PopString(); 
	return w
}

func (w *GuildChatWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.MemberID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.IconID) 
	pck.PutInt64(w.ChatTime) 
	pck.PutString(w.Msg) 
}

