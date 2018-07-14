package wx

// MsgType	说明
// 1	文本消息
// 3	图片消息
// 34	语音消息
// 37	好友确认消息
// 40	POSSIBLEFRIEND_MSG
// 42	共享名片
// 43	视频消息
// 47	动画表情
// 48	位置消息
// 49	分享链接
// 50	VOIPMSG
// 51	微信初始化消息
// 52	VOIPNOTIFY
// 53	VOIPINVITE
// 62	小视频
// 9999	SYSNOTICE
// 10000	系统消息
// 10002	撤回消息
const (
	MSGTYPE_TEXT               = 1
	MSGTYPE_IMAGE              = 3
	MSGTYPE_VOICE              = 34
	MSGTYPE_VIDEO              = 43
	MSGTYPE_MICROVIDEO         = 62
	MSGTYPE_EMOTICON           = 47
	MSGTYPE_APP                = 49
	MSGTYPE_VOIPMSG            = 50
	MSGTYPE_VOIPNOTIFY         = 52
	MSGTYPE_VOIPINVITE         = 53
	MSGTYPE_LOCATION           = 48
	MSGTYPE_STATUSNOTIFY       = 51
	MSGTYPE_SYSNOTICE          = 9999
	MSGTYPE_POSSIBLEFRIEND_MSG = 40
	MSGTYPE_VERIFYMSG          = 37
	MSGTYPE_SHARECARD          = 42
	MSGTYPE_SYS                = 10000
	MSGTYPE_RECALLED           = 10002
	//获取群组次数
	GET_GROUP_MEMBERS_TIMES = 6
)
