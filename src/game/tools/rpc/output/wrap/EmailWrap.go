package wraps

import (
	rpc "tipcat.com/common/rpc/simple"
)

type EmailWrap struct {
	
	ID int32 // 邮件id
	Title string // 邮件标题
	Content string // 邮件内容
	Flag int16 // 邮件是否读取
	Attachment string // 邮件附件
	AttachmentState int16 // 邮件附件是否已经领取
	EmailTime int64 // 邮件创建时间
	From string // 邮件发送者
	rpc.Wrapper
}

func (w *EmailWrap)Decode(pck *rpc.Packet) rpc.Wrapper {
	
	w.ID = pck.PopInt32(); 
	w.Title = pck.PopString(); 
	w.Content = pck.PopString(); 
	w.Flag = pck.PopInt16(); 
	w.Attachment = pck.PopString(); 
	w.AttachmentState = pck.PopInt16(); 
	w.EmailTime = pck.PopInt64(); 
	w.From = pck.PopString(); 
	return w
}

func (w *EmailWrap)Encode(pck *rpc.Packet) {
	
	pck.PutInt32(w.ID) 
	pck.PutString(w.Title) 
	pck.PutString(w.Content) 
	pck.PutInt16(w.Flag) 
	pck.PutString(w.Attachment) 
	pck.PutInt16(w.AttachmentState) 
	pck.PutInt64(w.EmailTime) 
	pck.PutString(w.From) 
}

