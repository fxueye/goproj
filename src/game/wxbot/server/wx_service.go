package server

import (
	"game/common/server/wx"

	log "github.com/cihub/seelog"
)

type WxService struct {
	wx.IMessgeHandler
	*wx.WxService
}

func newWxService(qrcodeDir string) *WxService {
	s := new(WxService)
	s.WxService = wx.NewWxService(qrcodeDir, s)
	return s
}

func (s *WxService) OnMessage(m *wx.Message) {
	if m.MsgType == 1 { // 文本消息
		log.Infof("%s: %s", s.GetUserName(m.FromUserName), m.Content)
	} else if m.MsgType == 3 { // 图片消息
	} else if m.MsgType == 34 { // 语音消息
	} else if m.MsgType == 43 { // 表情消息
	} else if m.MsgType == 47 { // 表情消息
	} else if m.MsgType == 49 { // 链接消息
	} else if m.MsgType == 51 { // 用户在手机进入某个联系人聊天界面时收到的消息
	} else {
		log.Infof("%s: MsgType: %d", s.GetUserName(m.FromUserName), m.MsgType)
	}
}
