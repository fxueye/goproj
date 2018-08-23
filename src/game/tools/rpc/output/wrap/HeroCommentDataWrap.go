package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type HeroCommentDataWrap struct {
	
	CommentID int32 // 评论ID
	HeroID int32 // 英雄ID
	UserID int64 // 用户ID
	UserNick string // 用户昵称
	Content string // 内容
	LikeNum int32 // 点赞数
	CommentTime int64 // 评论时间
	IsLiked bool // 是否已经点赞
	rpc.Wrapper
}

func (w *HeroCommentDataWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.CommentID = pck.PopInt32(); 
	w.HeroID = pck.PopInt32(); 
	w.UserID = pck.PopInt64(); 
	w.UserNick = pck.PopString(); 
	w.Content = pck.PopString(); 
	w.LikeNum = pck.PopInt32(); 
	w.CommentTime = pck.PopInt64(); 
	w.IsLiked = pck.PopBool(); 
	return w
}

func (w *HeroCommentDataWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.CommentID) 
	pck.PutInt32(w.HeroID) 
	pck.PutInt64(w.UserID) 
	pck.PutString(w.UserNick) 
	pck.PutString(w.Content) 
	pck.PutInt32(w.LikeNum) 
	pck.PutInt64(w.CommentTime) 
	pck.PutBool(w.IsLiked) 
}

