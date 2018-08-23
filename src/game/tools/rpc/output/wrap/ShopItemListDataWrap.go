package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ShopItemListDataWrap struct {
	
	ShopItemList []*ShopItemDataWrap // 商店item数据
	NextTime int64 // 下次跟新时间
	rpc.Wrapper
}

func (w *ShopItemListDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ShopItemList = make([]*ShopItemDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.ShopItemList); i++ {
		w.ShopItemList[i] = new(ShopItemDataWrap)
		w.ShopItemList[i].Decode(pck)
		
	}
	
	w.NextTime = pck.PopInt64(); 
	return w
}

func (w *ShopItemListDataWrap)Encode(pck *rpc.Packet) {
	
	if w.ShopItemList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.ShopItemList)))
    	for i := 0; i < len(w.ShopItemList); i++ {
    		w.ShopItemList[i].Encode(pck);
			
    	}
    }
	
	pck.PutInt64(w.NextTime) 
}

