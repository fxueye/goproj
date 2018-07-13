package server

import (
	"game/common/server/wx"

	log "github.com/cihub/seelog"
)

type WxService struct {
	wx.IMessgeHandler
	*wx.WxService
}

func newWxService(loginUrl, qrcodeDir string) *WxService {
	s := new(WxService)
	s.WxService = wx.NewWxService(loginUrl, qrcodeDir, config.Special, s)
	return s
}

func (s *WxService) OnMessage(m *wx.Message) {
	log.Infof("%v", *m)
	if m.MsgType == wx.MSGTYPE_TEXT { // 文本消息
		for _, nickName := range config.ForwardUserNames {
			user, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.SendMsg(user.UserName, m.Content)
		}
		log.Infof("%s: %s", s.GetNickName(m.FromUserName), m.Content)
	} else if m.MsgType == wx.MSGTYPE_IMAGE { // 图片消息
	} else if m.MsgType == wx.MSGTYPE_VOICE { // 语音消息
	} else if m.MsgType == wx.MSGTYPE_VIDEO { // 表情消息
	} else if m.MsgType == wx.MSGTYPE_EMOTICON { // 表情消息
	} else if m.MsgType == wx.MSGTYPE_APP { // 链接消息
	} else if m.MsgType == wx.MSGTYPE_STATUSNOTIFY { // 用户在手机进入某个联系人聊天界面时收到的消息
	} else {
		log.Infof("%s: MsgType: %d", s.GetNickName(m.FromUserName), m.MsgType)
	}
	// s.SendMsg(m.FromUserName, "您好！有什么能为您效劳的？")
}
