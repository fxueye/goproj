package wraps

import (
	rpc "game/common/rpc/simple"
)

type PlayerWrap struct {
	
	UID int64 // 游戏的uid
	GUID string // 平台唯一ID,游客登录的话为设备ID
	PID string // 登录平台ID
	UserName string // 用户名（唯一，初始和UID一致）
	LoginTime int64 // 登录时间戳（秒）
	OnlineTime int64 // 总在线时长（秒）
	CreateTime int64 // 创角时间戳（秒）
	ServerTime int64 // 服务器时间（秒）
	OnBattleIdx int32 // 上阵索引
	Items []*ItemDataWrap // 道具列表
	Icon int32 // 头像
	TutorialMask int64 // 新手引导
	rpc.Wrapper
}

func (w *PlayerWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.UID = pck.PopInt64(); 
	w.GUID = pck.PopString(); 
	w.PID = pck.PopString(); 
	w.UserName = pck.PopString(); 
	w.LoginTime = pck.PopInt64(); 
	w.OnlineTime = pck.PopInt64(); 
	w.CreateTime = pck.PopInt64(); 
	w.ServerTime = pck.PopInt64(); 
	w.OnBattleIdx = pck.PopInt32(); 
	w.Items = make([]*ItemDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Items); i++ {
		w.Items[i] = new(ItemDataWrap)
		w.Items[i].Decode(pck)
		
	}
	
	w.Icon = pck.PopInt32(); 
	w.TutorialMask = pck.PopInt64(); 
	return w
}

func (w *PlayerWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt64(w.UID) 
	pck.PutString(w.GUID) 
	pck.PutString(w.PID) 
	pck.PutString(w.UserName) 
	pck.PutInt64(w.LoginTime) 
	pck.PutInt64(w.OnlineTime) 
	pck.PutInt64(w.CreateTime) 
	pck.PutInt64(w.ServerTime) 
	pck.PutInt32(w.OnBattleIdx) 
	if w.Items == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Items)))
    	for i := 0; i < len(w.Items); i++ {
    		w.Items[i].Encode(pck);
			
    	}
    }
	
	pck.PutInt32(w.Icon) 
	pck.PutInt64(w.TutorialMask) 
}

