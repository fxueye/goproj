package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildSyncWrap struct {
	
	GuildID int32 // 公会ID
	GuildName string // 名字
	GuildIcon int32 // 公会图标
	rpc.Wrapper
}

func (w *GuildSyncWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GuildID = pck.PopInt32(); 
	w.GuildName = pck.PopString(); 
	w.GuildIcon = pck.PopInt32(); 
	return w
}

func (w *GuildSyncWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GuildID) 
	pck.PutString(w.GuildName) 
	pck.PutInt32(w.GuildIcon) 
}

