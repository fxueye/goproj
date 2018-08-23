package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ArenaWrap struct {
	
	Heroes []*BattleHeroWrap // 可选英雄
	OnBattleHeroes []*BattleHeroWrap // 上阵英雄
	rpc.Wrapper
}

func (w *ArenaWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.Heroes = make([]*BattleHeroWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.Heroes); i++ {
		w.Heroes[i] = new(BattleHeroWrap)
		w.Heroes[i].Decode(pck)
		
	}
	
	w.OnBattleHeroes = make([]*BattleHeroWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.OnBattleHeroes); i++ {
		w.OnBattleHeroes[i] = new(BattleHeroWrap)
		w.OnBattleHeroes[i].Decode(pck)
		
	}
	
	return w
}

func (w *ArenaWrap)Encode(pck *rpc.Packet) {
	
	if w.Heroes == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.Heroes)))
    	for i := 0; i < len(w.Heroes); i++ {
    		w.Heroes[i].Encode(pck);
			
    	}
    }
	
	if w.OnBattleHeroes == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.OnBattleHeroes)))
    	for i := 0; i < len(w.OnBattleHeroes); i++ {
    		w.OnBattleHeroes[i].Encode(pck);
			
    	}
    }
	
}

