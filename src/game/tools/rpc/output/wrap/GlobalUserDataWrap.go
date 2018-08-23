package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GlobalUserDataWrap struct {
	
	Uid int64 // did
	Username string // 用户名
	Icon int32 // 头像
	Score int32 // 积分
	Level int32 // 等级
	OnlineTime int64 // 在线时间
	SessID int64 // session
	Reconnect bool // 是否重连
	rpc.Wrapper
}

func (w *GlobalUserDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Uid = pck.PopInt64(); 
	w.Username = pck.PopString(); 
	w.Icon = pck.PopInt32(); 
	w.Score = pck.PopInt32(); 
	w.Level = pck.PopInt32(); 
	w.OnlineTime = pck.PopInt64(); 
	w.SessID = pck.PopInt64(); 
	w.Reconnect = pck.PopBool(); 
	return w
}

func (w *GlobalUserDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.Uid) 
	pck.PutString(w.Username) 
	pck.PutInt32(w.Icon) 
	pck.PutInt32(w.Score) 
	pck.PutInt32(w.Level) 
	pck.PutInt64(w.OnlineTime) 
	pck.PutInt64(w.SessID) 
	pck.PutBool(w.Reconnect) 
}

