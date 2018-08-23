package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type MemberSyncWrap struct {
	
	UID int64 // 玩家UID
	GuildID int32 // 公会ID
	GuildName string // 名字
	GuildRank int16 // 公会官职
	GuildIcon int32 // 公会图标
	rpc.Wrapper
}

func (w *MemberSyncWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.GuildID = pck.PopInt32(); 
	w.GuildName = pck.PopString(); 
	w.GuildRank = pck.PopInt16(); 
	w.GuildIcon = pck.PopInt32(); 
	return w
}

func (w *MemberSyncWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutInt32(w.GuildID) 
	pck.PutString(w.GuildName) 
	pck.PutInt16(w.GuildRank) 
	pck.PutInt32(w.GuildIcon) 
}

