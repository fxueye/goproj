package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type GuildDetailWrap struct {
	
	GuildID int32 // 公会id
	GuildName string // 名字
	IconID int32 // 头像ID
	GuildScore int32 // 公会积分
	GuildNote string // 公会公告
	JoinType int16 // 加入类型
	JoinScore int32 // 加入积分
	TopRank int32 // 公会排名
	WeekDonate int32 // 公会本周捐献
	Members []*GuildMemberBriefWrap // 成员
	rpc.Wrapper
}

func (w *GuildDetailWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.GuildID = pck.PopInt32(); 
	w.GuildName = pck.PopString(); 
	w.IconID = pck.PopInt32(); 
	w.GuildScore = pck.PopInt32(); 
	w.GuildNote = pck.PopString(); 
	w.JoinType = pck.PopInt16(); 
	w.JoinScore = pck.PopInt32(); 
	w.TopRank = pck.PopInt32(); 
	w.WeekDonate = pck.PopInt32(); 
	w.Members = make([]*GuildMemberBriefWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Members); i++ {
		w.Members[i] = new(GuildMemberBriefWrap)
		w.Members[i].Decode(pck)
		
	}
	
	return w
}

func (w *GuildDetailWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.GuildID) 
	pck.PutString(w.GuildName) 
	pck.PutInt32(w.IconID) 
	pck.PutInt32(w.GuildScore) 
	pck.PutString(w.GuildNote) 
	pck.PutInt16(w.JoinType) 
	pck.PutInt32(w.JoinScore) 
	pck.PutInt32(w.TopRank) 
	pck.PutInt32(w.WeekDonate) 
	if w.Members == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Members)))
    	for i := 0; i < len(w.Members); i++ {
    		w.Members[i].Encode(pck);
			
    	}
    }
	
}

