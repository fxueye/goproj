package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type HeroCommentListDataWrap struct {
	
	CommentsList []*HeroCommentDataWrap // 评论列表
	rpc.Wrapper
}

func (w *HeroCommentListDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.CommentsList = make([]*HeroCommentDataWrap, int(pck.PopInt16()))
	for i := 0; i < len(w.CommentsList); i++ {
		w.CommentsList[i] = new(HeroCommentDataWrap)
		w.CommentsList[i].Decode(pck)
		
	}
	
	return w
}

func (w *HeroCommentListDataWrap)Encode(pck *rpc.Packet) {
	
	if w.CommentsList == nil {
		pck.PutInt16(0)
	} else {
    	pck.PutInt16(int16(len(w.CommentsList)))
    	for i := 0; i < len(w.CommentsList); i++ {
    		w.CommentsList[i].Encode(pck);
			
    	}
    }
	
}

