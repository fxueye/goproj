package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ShopItemDataWrap struct {
	
	ShopItemID int32 // 商品道具ID
	GoodID int32 // 商品id 钻石 金币 宝箱 随机英雄 固定英雄
	GoodType int32 // 商品类型 1 买钻石 2 买金币 3 买宝箱 4 买随机英雄 5 买固定英雄
	GoodCount int32 // 商品数量
	BuyTime int64 // 购买时间
	Count int32 // 商品购买数量
	rpc.Wrapper
}

func (w *ShopItemDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ShopItemID = pck.PopInt32(); 
	w.GoodID = pck.PopInt32(); 
	w.GoodType = pck.PopInt32(); 
	w.GoodCount = pck.PopInt32(); 
	w.BuyTime = pck.PopInt64(); 
	w.Count = pck.PopInt32(); 
	return w
}

func (w *ShopItemDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ShopItemID) 
	pck.PutInt32(w.GoodID) 
	pck.PutInt32(w.GoodType) 
	pck.PutInt32(w.GoodCount) 
	pck.PutInt64(w.BuyTime) 
	pck.PutInt32(w.Count) 
}

