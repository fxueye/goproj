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
func (s *WxService) OnWxInitSucces() {
	if len(s.receptionist) == 0 {
		for _, nickName := range config.ForwardUserNames {
			user, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.receptionist[user.UserName] = user
			msg := fmt.Sprintf(config.TextConfig[18])
			s.SendMsg(user.UserName, msg)
		}
	}
}
func (s *WxService) OnMessage(m *wx.Message) {

	if s.handlerReceptionMsg(m) {
		return
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
	if s.handlerGuest(m) {
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
			content = fmt.Sprintf(config.TextConfig[0], s.ClearCharactert(group.NickName), s.ClearCharactert(sendUser.NickName), content)
		} else {
			content = fmt.Sprintf(config.TextConfig[1], s.ClearCharactert(friendNickName), s.ClearCharactert(content))
		}

	} else if m.MsgType == wx.MSGTYPE_IMAGE { // 图片消息
		if isGroupMsg {
			content = fmt.Sprintf(config.TextConfig[2], s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf(config.TextConfig[3], s.ClearCharactert(friendNickName))
		}
	} else if m.MsgType == wx.MSGTYPE_VOICE { // 语音消息
		if isGroupMsg {
			content = fmt.Sprintf(config.TextConfig[4], s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf(config.TextConfig[5], s.ClearCharactert(friendNickName))
		}
	} else if m.MsgType == wx.MSGTYPE_VIDEO { // 表情消息
		if isGroupMsg {
			content = fmt.Sprintf(config.TextConfig[6], s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf(config.TextConfig[7], s.ClearCharactert(friendNickName))
		}

	} else if m.MsgType == wx.MSGTYPE_EMOTICON { // 表情消息
		if isGroupMsg {
			content = fmt.Sprintf(config.TextConfig[8], s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf(config.TextConfig[9], s.ClearCharactert(friendNickName))
		}

	} else if m.MsgType == wx.MSGTYPE_APP { // 链接消息
		if isGroupMsg {
			content = fmt.Sprintf(config.TextConfig[10], s.ClearCharactert(friendNickName))
		} else {
			content = fmt.Sprintf(config.TextConfig[11], s.ClearCharactert(friendNickName))
		}

	} else if m.MsgType == wx.MSGTYPE_STATUSNOTIFY { // 用户在手机进入某个联系人聊天界面时收到的消息
		if isGroupMsg {
			content = fmt.Sprintf(config.TextConfig[12])
		} else {
			content = fmt.Sprintf(config.TextConfig[13])
		}

	} else {
		log.Infof("%s: MsgType: %d", s.GetNickName(m.FromUserName), m.MsgType)
	}
	for _, user := range s.receptionist {
		s.SendMsg(user.UserName, content)
	}
	// s.SendMsg(m.FromUserName, "您好！有什么能为您效劳的？")
}
func (s *WxService) handlerGuest(m *wx.Message) bool {
	guest, err := s.GetUser(m.FromUserName)
	if err != nil {
		return false
	}
	if guest.Rec != nil {
		s.SendMsg(guest.Rec.UserName, m.Content)
		return true
	}
	return false
}
func (s *WxService) handlerReceptionMsg(m *wx.Message) bool {
	if _, ok := s.receptionist[m.FromUserName]; ok {
		recUser := s.receptionist[m.FromUserName]
		content := m.Content
		strs := strings.Split(content, ":")
		if len(strs) > 0 {
			switch strs[0] {
			case "a": //添加接待员
				nickName := strs[1]
				user, err := s.GetUserByNickName(nickName)
				if err != nil {
					log.Infof("receptionist not find:%s", nickName)
					return true
				}
				if _, ok := s.receptionist[user.UserName]; ok {
					log.Infof("receptionist is exist:%s", nickName)
					return true
				}
				s.receptionist[user.UserName] = user
				msg := fmt.Sprintf(config.TextConfig[14], nickName)
				s.SendMsg(recUser.UserName, msg)
			case "r": //接待XXX
				nickName := strs[1]
				user, err := s.GetUserByNickName(nickName)
				if err != nil {
					log.Infof("receptionist not find:%s", nickName)
					return true
				}
				if user.Rec != nil {
					s.SendMsg(recUser.UserName, config.TextConfig[20])
					return true
				}
				user.Rec = recUser
				recUser.Rec = user
				msg := fmt.Sprintf(config.TextConfig[15], nickName)
				s.SendMsg(recUser.UserName, msg)
				for _, u := range s.receptionist {
					if recUser.UserName != u.UserName {
						msg := fmt.Sprintf(config.TextConfig[17], recUser.NickName, user.NickName)
						s.SendMsg(u.UserName, msg)
					}
				}
			case "e": //结束接待
				user := recUser.Rec
				if user == nil {
					msg := fmt.Sprintf(config.TextConfig[19])
					s.SendMsg(recUser.UserName, msg)
					return true
				}
				msg := fmt.Sprintf(config.TextConfig[16], user.NickName)
				s.SendMsg(recUser.UserName, msg)
				user.Rec = nil
				recUser.Rec = nil

			case "h": //帮助
				msg := fmt.Sprintf(config.TextConfig[18])
				s.SendMsg(recUser.UserName, msg)
			default:
				if recUser.Rec != nil {
					s.SendMsg(recUser.Rec.UserName, content)
				}
			}

		}
		return true
	}
	return false
}
func (s *WxService) ClearCharactert(str string) string {
	str = strings.Replace(str, "<br/>", "\n", -1)
	str = strings.Replace(str, "&amp;", "&", -1)
	exp := regexp.MustCompile(`<span class=".*?"></span>`)
	str = exp.ReplaceAllString(str, "")
	return str
}
