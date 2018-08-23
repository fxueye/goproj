package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type BattleReportWrap struct {
	
	Version int32 // 战报版本号
	ReportID string // 战报ID，服务端生成
	BattleType int16 // 战斗类型（0 天梯，1 竞技场，2 练习赛，3 关卡）
	DungeonID int32 // 活动副本ID(战斗类型为关卡时使用）
	QuestID int32 // 关卡ID(战斗类型为关卡时使用）
	Seed int64 // 随机数种子
	Timestamp int64 // 时间戳
	FirstCamp int16 // 先手阵营
	Result int16 // 战斗结果(1 A赢，2 B赢，3 平局，4 A掉线，5 B掉线）
	CampA *BattlePlayerWrap // 阵营A
	CampB *BattlePlayerWrap // 阵营B（战斗类型不为关卡时使用）
	Cmds []*BattleCmdWrap // 指令列表
	rpc.Wrapper
}

func (w *BattleReportWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Version = pck.PopInt32(); 
	w.ReportID = pck.PopString(); 
	w.BattleType = pck.PopInt16(); 
	w.DungeonID = pck.PopInt32(); 
	w.QuestID = pck.PopInt32(); 
	w.Seed = pck.PopInt64(); 
	w.Timestamp = pck.PopInt64(); 
	w.FirstCamp = pck.PopInt16(); 
	w.Result = pck.PopInt16(); 
	w.CampA = new(BattlePlayerWrap)
	w.CampA.Decode(pck)
	
	w.CampB = new(BattlePlayerWrap)
	w.CampB.Decode(pck)
	
	w.Cmds = make([]*BattleCmdWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Cmds); i++ {
		w.Cmds[i] = new(BattleCmdWrap)
		w.Cmds[i].Decode(pck)
		
	}
	
	return w
}

func (w *BattleReportWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.Version) 
	pck.PutString(w.ReportID) 
	pck.PutInt16(w.BattleType) 
	pck.PutInt32(w.DungeonID) 
	pck.PutInt32(w.QuestID) 
	pck.PutInt64(w.Seed) 
	pck.PutInt64(w.Timestamp) 
	pck.PutInt16(w.FirstCamp) 
	pck.PutInt16(w.Result) 
	w.CampA.Encode(pck); 
	w.CampB.Encode(pck); 
	if w.Cmds == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Cmds)))
    	for i := 0; i < len(w.Cmds); i++ {
    		w.Cmds[i].Encode(pck);
			
    	}
    }
	
}

