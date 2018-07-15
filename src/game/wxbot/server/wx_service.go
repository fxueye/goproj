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
	//写入存储开启取出
	datas map[string][]*Stroke
}

func newWxService(loginUrl, qrcodeDir string) *WxService {
	s := new(WxService)
	s.WxService = wx.NewWxService(loginUrl, qrcodeDir, config.Special, s)
	s.datas = make(map[string][]*Stroke)
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

			stroke := new(Stroke)
			stroke.Send = s.ClearCharactert(sendUser.NickName)
			iphones := utils.GetTelNum(content)
			iphones = utils.RemoveDuplicatesAndEmpty(iphones)
			stroke.Tel = strings.Join(iphones, ",")
			stroke.Content = content
			stroke.Timestamp = utils.NowTimestamp()
			if _, ok := s.datas[stroke.Send]; ok {
				strokes := s.datas[stroke.Send]
				for _, s := range strokes {
					if s.Content == stroke.Content {
						return
					}
				}
			}

			s.datas[stroke.Send] = append(s.datas[stroke.Send], stroke)
			err = CreateStroke(*stroke)
			if err != nil {
				log.Error(err)
			}
			//消息转发
			for _, nickName := range config.ForwardUserNames {
				user, err := s.GetUserByNickName(nickName)
				if err != nil {
					continue
				}
				content = fmt.Sprintf("群:[%s]:[%s]:\n%s", s.ClearCharactert(group.NickName), s.ClearCharactert(sendUser.NickName), content)
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

				content = fmt.Sprintf("好友:%s:\n%s", s.ClearCharactert(friendNickName), content)
				s.SendMsg(user.UserName, content)
			}
		}

	} else if m.MsgType == wx.MSGTYPE_IMAGE { // 图片消息
		friendNickName := s.GetNickName(m.FromUserName)
		content := fmt.Sprintf("好友:%s:\n 图片消息", s.ClearCharactert(friendNickName))
		for _, nickName := range config.ForwardUserNames {
			user, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.SendMsg(user.UserName, content)
		}
	} else if m.MsgType == wx.MSGTYPE_VOICE { // 语音消息
		friendNickName := s.GetNickName(m.FromUserName)
		content := fmt.Sprintf("好友:%s:\n 语音消息", s.ClearCharactert(friendNickName))
		for _, nickName := range config.ForwardUserNames {
			user, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.SendMsg(user.UserName, content)
		}
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
