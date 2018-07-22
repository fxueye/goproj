package server

import (
	"fmt"
	"game/common/server/wx"
	"game/common/utils"
	"regexp"
	"strings"

	log "github.com/cihub/seelog"
)

type WxService struct {
	wx.IMessgeHandler
	*wx.WxService
	receptionist map[string]*wx.User
}

func newWxService(loginUrl, qrcodeDir string) *WxService {
	s := new(WxService)
	s.WxService = wx.NewWxService(loginUrl, qrcodeDir, config.Special, s)
	s.receptionist = make(map[string]*wx.User)
	return s
}

func (s *WxService) OnMessage(m *wx.Message) {

	if len(s.receptionist) == 0 {
		for _, nickName := range config.ForwardUserNames {
			user, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.receptionist[nickName] = user
		}
	}

	formUserName := m.FromUserName
	isGroupMsg := strings.Index(formUserName, "@@") != -1
	if isGroupMsg && !config.GroupMsg {
		return
	}
	friendNickName := s.GetNickName(m.FromUserName)
	if utils.IsInStringArray(config.Special, friendNickName) {
		log.Infof("特殊帐号消息: %v", friendNickName)
		return
	}
	content := m.Content
	if m.MsgType == wx.MSGTYPE_TEXT { // 文本消息
		if isGroupMsg {
			group, err := s.GetGroup(formUserName)

			if err != nil {
				log.Errorf("%v", err)
				return
			}

			index := strings.Index(m.Content, ":")
			sendUserName := string([]byte(m.Content)[0:index])
			content = string([]byte(m.Content)[index+1:])
			content = s.ClearCharactert(content)
			sendUser, err := s.GetGroupMember(formUserName, sendUserName)
			if err != nil {
				log.Errorf("%v", err)
				return
			}
			//消息转发
			content = fmt.Sprintf("群:[%s]:[%s]:\n%s", s.ClearCharactert(group.NickName), s.ClearCharactert(sendUser.NickName), content)
		} else {
			content = fmt.Sprintf("好友:%s:\n%s", s.ClearCharactert(friendNickName), s.ClearCharactert(content))
		}

	} else if m.MsgType == wx.MSGTYPE_IMAGE { // 图片消息
		if isGroupMsg {
			content = fmt.Sprintf("好友:%s:\n 图片消息", s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf("好友:%s:\n 图片消息", s.ClearCharactert(friendNickName))
		}
	} else if m.MsgType == wx.MSGTYPE_VOICE { // 语音消息
		if isGroupMsg {
			content = fmt.Sprintf("好友:%s:\n 语音消息", s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf("好友:%s:\n 语音消息", s.ClearCharactert(friendNickName))
		}
	} else if m.MsgType == wx.MSGTYPE_VIDEO { // 表情消息
		if isGroupMsg {
			content = fmt.Sprintf("好友:%s:\n 表情消息", s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf("好友:%s:\n 表情消息", s.ClearCharactert(friendNickName))
		}

	} else if m.MsgType == wx.MSGTYPE_EMOTICON { // 表情消息
		if isGroupMsg {
			content = fmt.Sprintf("好友:%s:\n 表情消息", s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf("好友:%s:\n 表情消息", s.ClearCharactert(friendNickName))
		}

	} else if m.MsgType == wx.MSGTYPE_APP { // 链接消息
		if isGroupMsg {
			content = fmt.Sprintf("好友:%s:\n 链接消息", s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf("好友:%s:\n 链接消息", s.ClearCharactert(friendNickName))
		}

	} else if m.MsgType == wx.MSGTYPE_STATUSNOTIFY { // 用户在手机进入某个联系人聊天界面时收到的消息
		if isGroupMsg {
			content = fmt.Sprintf("小手机有人在使用微信!\n")
		} else {
			content = fmt.Sprintf("小手机有人在使用微信!\n")
		}

	} else {
		log.Infof("%s: MsgType: %d", s.GetNickName(m.FromUserName), m.MsgType)
	}
	for _, user := range s.receptionist {
		s.SendMsg(user.UserName, content)
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
