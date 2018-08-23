package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildMemberBriefWrap struct {
	
	UID int64 // 玩家UID
	Name string // 名字
	Level int32 // 等级
	Score int32 // 积分
	GuildRank int16 // 公会官职
	Donation int32 // 捐献数
	Online bool // 是否在线
	Icon int32 // 头像
	rpc.Wrapper
}

func (w *GuildMemberBriefWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.Name = pck.PopString(); 
	w.Level = pck.PopInt32(); 
	w.Score = pck.PopInt32(); 
	w.GuildRank = pck.PopInt16(); 
	w.Donation = pck.PopInt32(); 
	w.Online = pck.PopBool(); 
	w.Icon = pck.PopInt32(); 
	return w
}

func (w *GuildMemberBriefWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutString(w.Name) 
	pck.PutInt32(w.Level) 
	pck.PutInt32(w.Score) 
	pck.PutInt16(w.GuildRank) 
	pck.PutInt32(w.Donation) 
	pck.PutBool(w.Online) 
	pck.PutInt32(w.Icon) 
}

