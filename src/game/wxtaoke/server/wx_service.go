package server

import (
	"encoding/json"
	"fmt"
	"game/common/server/wx"
	"game/common/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

type WxService struct {
	wx.IMessgeHandler
	*wx.WxService
	controller map[string]*wx.User
	proGroups  map[string]*wx.User
	timers     []int64
}

func newWxService(loginUrl, qrcodeDir, tempImgDir string) *WxService {
	s := new(WxService)
	s.WxService = wx.NewWxService(loginUrl, qrcodeDir, tempImgDir, config.Special, s)
	s.controller = make(map[string]*wx.User)
	s.proGroups = make(map[string]*wx.User)
	s.timers = make([]int64, 0)
	return s
}

func (s *WxService) OnWxInitSucces() {
	if len(s.controller) == 0 {
		for _, nickName := range config.ControllerUserNames {
			user, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.controller[user.UserName] = user
			msg := fmt.Sprintf(config.TextConfig[18])
			s.SendMsg(user.UserName, msg)
		}
	}
	if len(s.proGroups) == 0 {
		for _, nickName := range config.Groups {
			group, err := s.GetUserByNickName(nickName)
			if err != nil {
				continue
			}
			s.proGroups[group.UserName] = group
		}
	}
	go s.startTimer()
}
func (s *WxService) OnMessage(m *wx.Message) {
	if s.handlerReceptionMsg(m) {
		return
	}

	formUserName := m.FromUserName
	isGroupMsg := strings.Index(formUserName, "@@") != -1
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
			index = strings.Index(content, "@")
			if index == 1 {
				indexT := strings.Index(content, " ")
				name := ""
				keyword := ""
				if indexT == -1 {
					name = string([]byte(content)[index+1:])
				} else {
					name = string([]byte(content)[index+1 : indexT])
					keyword = string([]byte(content)[indexT+1:])
				}
				if name == s.LoginUser().NickName {
					couMap := GetCoupon(keyword)
					coupon := MakeCouponStr(couMap)
					s.SendMsg(formUserName, coupon)
				} else {
					log.Infof("content:%v", content)
					return
				}

			}
			content = fmt.Sprintf(config.TextConfig[0], s.ClearCharactert(group.NickName), s.ClearCharactert(sendUser.NickName), content)
			// s.SendMsg(formUserName, content)
		} else {
			content = fmt.Sprintf(config.TextConfig[1], s.ClearCharactert(friendNickName), s.ClearCharactert(content))
		}
		log.Infof("content:%v", content)

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
	// for _, user := range s.controller {
	// s.SendMsg(user.UserName, content)
	// }
	// for _, group := range s.proGroups {
	// 	s.SendMsg(user.UserName, content)
	// }
	// s.SendMsg(m.FromUserName, "您好！有什么能为您效劳的？")
}

func (s *WxService) sendProGroups(msg string) {
	for _, group := range s.proGroups {
		s.SendMsg(group.UserName, msg)
	}
}
func (s *WxService) sendProGroupsImg(path string) {
	for _, group := range s.proGroups {
		err := s.SendImg(group.UserName, path)
		if err != nil {
			continue
		}
	}
	os.Remove(path)
}
func (s *WxService) sendController(msg string) {
	for _, user := range s.controller {
		s.SendMsg(user.UserName, msg)
	}
}

func (s *WxService) handlerReceptionMsg(m *wx.Message) bool {
	recUser := s.controller[m.FromUserName]
	content := m.Content
	strs := strings.Split(content, ":")
	if len(strs) > 0 {
		switch strs[0] {
		case "add user": //添加
			nickName := strs[1]
			user := recUser
			var err error
			if nickName != "" {
				user, err = s.GetUserByNickName(nickName)
				if err != nil {
					log.Infof("controller not find:%s", nickName)
					return true
				}
			}
			if _, ok := s.controller[user.UserName]; ok {
				log.Infof("controller is exist:%s", nickName)
				return true
			}
			s.controller[user.UserName] = user
			msg := fmt.Sprintf(config.TextConfig[14], nickName)
			s.SendMsg(recUser.UserName, msg)
			return true
		case "add group":
			nickName := strs[1]
			if nickName != "" {
				group, err := s.GetUserByNickName(nickName)
				if err != nil {
					log.Infof("group not find:%s", nickName)
					return true
				}
				s.proGroups[group.UserName] = group
				s.SendMsg(recUser.UserName, "add success!")
			} else {
				s.SendMsg(recUser.UserName, "comd is add group:groupname")
				log.Info("comd is add group:groupname")
			}
		case "get users":
			users := s.GetUsers()
			msg := "好友:\n"
			for _, v := range users {
				msg += fmt.Sprintf("[%v]\n", v.NickName)
			}
			s.SendMsg(recUser.UserName, msg)
			log.Infof("msg:%v", msg)
			return true
		case "get groups":
			groups := s.GetGroups()
			msg := "群组:\n"
			for _, v := range groups {
				msg += fmt.Sprintf("[%v]\n", v.NickName)
			}
			s.SendMsg(recUser.UserName, msg)
			log.Infof("msg:%v", msg)
			return true
		case "send cou":
			msg := strs[1]
			s.sendProGroups(msg)
			return true
		case "get cou":
			couMap := GetCoupon("")
			data := couMap["data"].(map[string]interface{})
			imgStr := data["small_images"].(string)
			var smallImages []string
			err := json.Unmarshal([]byte(imgStr), &smallImages)
			path, err := s.GetImg(smallImages[0])
			if err != nil {
				log.Error(err)
				return false
			}
			err = s.SendImg(recUser.UserName, path)
			if err == nil {
				os.Remove(path)
			}
			coupon := MakeCouponStr(couMap)
			s.SendMsg(recUser.UserName, coupon)
			return true
		case "add timer":
			t := strs[1]
			ss := strings.Split(t, " ")
			if len(ss) == 2 {
				h, _ := strconv.Atoi(ss[0])
				i, _ := strconv.Atoi(ss[1])
				day := time.Now().Format("2006-01-02")
				timestr := fmt.Sprintf("%s %02d:%02d:00", day, h, i)
				t, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
				s.timers = append(s.timers, t.Unix())
				s.SendMsg(recUser.UserName, "add success!")
			} else {
				s.SendMsg(recUser.UserName, "add failed timer type error!")
			}
		case "h": //帮助
			msg := fmt.Sprintf(config.TextConfig[18])
			s.SendMsg(recUser.UserName, msg)
			return true
		default:
			return false
		}

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
func (s *WxService) startTimer() {
	s.initTimers()
	timer := time.NewTicker(time.Millisecond * 200)
	for {
		select {
		case <-timer.C:
			s.timerCheck()
		}
	}
}
func (s *WxService) initTimers() {
	for _, t := range config.SendTimer {
		ss := strings.Split(t, " ")
		if len(ss) == 2 {
			h, _ := strconv.Atoi(ss[0])
			i, _ := strconv.Atoi(ss[1])
			day := time.Now().Format("2006-01-02")
			timestr := fmt.Sprintf("%s %02d:%02d:00", day, h, i)
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
			s.timers = append(s.timers, t.Unix())
		}
	}
}
func (s *WxService) timerCheck() {
	for k, v := range s.timers {
		now := time.Now()
		if v == now.Unix() {
			s.SendCouToGroup()
			s.timers[k] = s.nextTime(v)
		} else if now.Unix() > v {
			s.timers[k] = s.nextTime(v)
		}
	}
}
func (s *WxService) nextTime(v int64) int64 {
	return time.Unix(v, 0).Add(time.Hour * 24).Unix()
}
func (s *WxService) SendCouToGroup() {
	couMap := GetCoupon("")
	data := couMap["data"].(map[string]interface{})
	imgStr := data["small_images"].(string)
	var smallImages []string
	err := json.Unmarshal([]byte(imgStr), &smallImages)
	path, err := s.GetImg(smallImages[0])
	if err != nil {
		log.Error(err)
		return
	}
	s.sendProGroupsImg(path)
	coupon := MakeCouponStr(couMap)
	s.sendProGroups(coupon)
	log.Infof("%v", coupon)
}
