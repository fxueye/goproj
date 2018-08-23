package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ProductInfoWrap struct {
	
	ProductID string // 商品ID
	SortID int16 // 排序ID
	ProductName string // 商品名
	Desc string // 描述
	Visible bool // 是否可见
	Hot bool // 是否热销
	MoneyType int32 // 货币类型
	MoneyCnt int32 // 发货数量
	GiftMoneyCnt int32 // 赠送数量
	Price float32 // 购买价格
	Currency string // 货币类型
	rpc.Wrapper
}

func (w *ProductInfoWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ProductID = pck.PopString(); 
	w.SortID = pck.PopInt16(); 
	w.ProductName = pck.PopString(); 
	w.Desc = pck.PopString(); 
	w.Visible = pck.PopBool(); 
	w.Hot = pck.PopBool(); 
	w.MoneyType = pck.PopInt32(); 
	w.MoneyCnt = pck.PopInt32(); 
	w.GiftMoneyCnt = pck.PopInt32(); 
	w.Price = pck.PopFloat32(); 
	w.Currency = pck.PopString(); 
	return w
}

func (w *ProductInfoWrap)Encode(pck *rpc.Packet) {
	
	pck.PutString(w.ProductID) 
	pck.PutInt16(w.SortID) 
	pck.PutString(w.ProductName) 
	pck.PutString(w.Desc) 
	pck.PutBool(w.Visible) 
	pck.PutBool(w.Hot) 
	pck.PutInt32(w.MoneyType) 
	pck.PutInt32(w.MoneyCnt) 
	pck.PutInt32(w.GiftMoneyCnt) 
	pck.PutFloat32(w.Price) 
	pck.PutString(w.Currency) 
}

