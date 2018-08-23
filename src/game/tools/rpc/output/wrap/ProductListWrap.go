package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type ProductListWrap struct {
	
	ProductList []*ProductInfoWrap // 商品列表
	rpc.Wrapper
}

func (w *ProductListWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ProductList = make([]*ProductInfoWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.ProductList); i++ {
		w.ProductList[i] = new(ProductInfoWrap)
		w.ProductList[i].Decode(pck)
		
	}
	
	return w
}

func (w *ProductListWrap)Encode(pck *rpc.Packet) {
	
	if w.ProductList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.ProductList)))
    	for i := 0; i < len(w.ProductList); i++ {
    		w.ProductList[i].Encode(pck);
			
    	}
    }
	
}

