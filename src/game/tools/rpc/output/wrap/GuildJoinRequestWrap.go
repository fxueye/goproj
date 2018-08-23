package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildJoinRequestWrap struct {
	
	UID int64 // 玩家UID
	Name string // 名字
	Level int32 // 等级
	Score int32 // 积分
	Icon int32 // 头像
	Msg string // 留言
	rpc.Wrapper
}

func (w *GuildJoinRequestWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.Name = pck.PopString(); 
	w.Level = pck.PopInt32(); 
	w.Score = pck.PopInt32(); 
	w.Icon = pck.PopInt32(); 
	w.Msg = pck.PopString(); 
	return w
}

func (w *GuildJoinRequestWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.Level) 
	pck.PutInt32(w.Score) 
	pck.PutInt32(w.Icon) 
	pck.PutString(w.Msg) 
}

