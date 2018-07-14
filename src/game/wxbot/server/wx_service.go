package server

import (
	"fmt"
	"game/common/server/wx"
	"regexp"
	"strings"

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
	if m.MsgType == wx.MSGTYPE_TEXT { // 文本消息
		formUserName := m.FromUserName
		if strings.Index(formUserName, "@@") != -1 {
			log.Infof("%s", formUserName)
			group, err := s.GetGroup(formUserName)

			if err != nil {
				log.Errorf("%v", err)
				return
			}

			index := strings.Index(m.Content, ":")
			sendUserName := string([]byte(m.Content)[0:index])
			content := string([]byte(m.Content)[index+1:])
			content = s.ClearCharactert(content)
			sendUser, err := s.GetGroupMember(formUserName, sendUserName)
			if err != nil {
				log.Errorf("%v", err)
				return
			}
			//消息转发
			for _, nickName := range config.ForwardUserNames {
				user, err := s.GetUserByNickName(nickName)
				if err != nil {
					continue
				}
				content = fmt.Sprintf("来自群:%s[%s]:\n%s", s.ClearCharactert(group.NickName), s.ClearCharactert(sendUser.NickName), content)
				s.SendMsg(user.UserName, content)
			}
		} else {
			for _, nickName := range config.ForwardUserNames {
				user, err := s.GetUserByNickName(nickName)
				if err != nil {
					continue
				}

				content := s.ClearCharactert(m.Content)
				friendNickName := s.GetNickName(m.FromUserName)

				content = fmt.Sprintf("来自好友:%s:\n%s", s.ClearCharactert(friendNickName), content)
				s.SendMsg(user.UserName, content)
			}
		}

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
func (s *WxService) ClearCharactert(str string) string {
	str = strings.Replace(str, "<br/>", "\n", -1)
	str = strings.Replace(str, "&amp;", "&", -1)
	exp := regexp.MustCompile(`<span class=".*?"></span>`)
	str = exp.ReplaceAllString(str, "")
	return str
}
