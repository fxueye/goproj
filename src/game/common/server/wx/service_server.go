package wx

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"game/common/server"
	"game/common/utils"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

var (
	fs                   = &FileStore{"/tmp/s.secret"}
	LoginUri             = "https://login.weixin.qq.com"
	ErrUnknow            = errors.New("Unknow Error")
	ErrUserNotExists     = errors.New("Error User Not Exist")
	ErrNotLogin          = errors.New("Not Login")
	ErrLoginTimeout      = errors.New("Login Timeout")
	ErrWaitingForConfirm = errors.New("Waiting For Confirm")
)

type syncStatus struct {
	Retcode  string
	Selector string
}

type WxService struct {
	server.BaseService

	httpClient  *Client
	secret      *wxSecret
	baseRequest *BaseRequest
	user        *User
	contacts    map[string]*User
	qrcode      string
	qrcodeDir   string
}

func NewWxService(qrcodeDir string) *WxService {
	s := new(WxService)
	s.httpClient = NewClient()
	s.secret = &wxSecret{}
	s.baseRequest = &BaseRequest{}
	s.user = &User{}
	s.contacts = make(map[string]*User)
	s.qrcodeDir = qrcodeDir
	return s
}

func (s *WxService) Start() error {
	s.BaseService.Start()
	newLoginUri, err := s.GetNewLoginUrl()
	if err != nil {
		return err
	}
	err = s.NewLoginPage(newLoginUri)
	if err != nil {
		return err
	}
	err = s.Init()
	if err != nil {
		return err
	}
	err = s.GetContacts()
	if err != nil {
		return err
	}
	s.AsyncDo(func() {
		defer func() {
			recover()
		}()
		err = s.Listening()
		if err != nil {
			return
		}
	})
	return nil
}

func (s *WxService) Init() error {
	values := &url.Values{}
	values.Set("r", TimestampStr())
	values.Set("lang", "en_US")
	values.Set("pass_ticket", s.secret.PassTicket)
	url := fmt.Sprintf("%s/webwxinit?%s", s.secret.BaseUri, values.Encode())
	s.baseRequest = &BaseRequest{
		Uin:      s.secret.Uin,
		Sid:      s.secret.Sid,
		Skey:     s.secret.Skey,
		DeviceID: s.secret.DeviceID,
	}
	b, err := s.httpClient.PostJson(url, map[string]interface{}{
		"BaseRequest": s.baseRequest,
	})
	if err != nil {
		log.Errorf("HTTP GET err: %s", err.Error())
		return err
	}
	var r InitResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	if r.BaseResponse.Ret == 0 {
		s.user = r.User
		s.updateSyncKey(r.SyncKey)
		return nil
	}
	return fmt.Errorf("Init error: %+v", r.BaseResponse)
}

func (s *WxService) Listening() error {
	err := s.TestSyncCheck()
	if err != nil {
		return err
	}
	for {
		syncStatus, err := s.SyncCheck()
		if err != nil {
			log.Errorf("sync check error: %s", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		if syncStatus.Retcode == "1100" {
			return errors.New("从微信客户端上登出")
		} else if syncStatus.Retcode == "1101" {
			return errors.New("从其它设备上登了网页微信")
		} else if syncStatus.Retcode == "0" {
			if syncStatus.Selector == "0" { // 无更新
				continue
			} else if syncStatus.Selector == "2" { // 有新消息
				ms, err := s.Sync()
				if err != nil {
					log.Errorf("sync err: %s", err.Error())
				}

				log.Info(ms)
				for _, m := range ms {
					s.HandleMsg(m)
				}
				// s.HandleMsgs(ms)
			} else { // 可能有其他类型的消息，直接丢弃
				log.Errorf("New Message, Unknow type: %+v", syncStatus)
				_, err := s.Sync()
				if err != nil {

				}
			}
		} else if syncStatus.Retcode == "1102" {
			return fmt.Errorf("Sync Error %+v", syncStatus)
		} else {
			log.Errorf("sync check Unknow Code: %+v", syncStatus)
		}
	}
}

func (wx *WxService) HandleMsg(m *Message) {
	if m.MsgType == 1 { // 文本消息
		log.Infof("%s: %s", wx.GetUserName(m.FromUserName), m.Content)
	} else if m.MsgType == 3 { // 图片消息
	} else if m.MsgType == 34 { // 语音消息
	} else if m.MsgType == 43 { // 表情消息
	} else if m.MsgType == 47 { // 表情消息
	} else if m.MsgType == 49 { // 链接消息
	} else if m.MsgType == 51 { // 用户在手机进入某个联系人聊天界面时收到的消息
	} else {
		log.Infof("%s: MsgType: %d", wx.GetUserName(m.FromUserName), m.MsgType)
	}
}
func (s *WxService) GetUserName(userName string) string {
	u, err := s.GetUser(userName)
	if err != nil {
		return userName
	}
	if u.RemarkName != "" {
		return u.RemarkName
	} else {
		return u.NickName
	}
}
func (wx *WxService) GetUser(userName string) (*User, error) {
	u, ok := wx.contacts[userName]
	if ok {
		return u, nil
	} else {
		return nil, ErrUserNotExists
	}
}
func (s *WxService) Sync() ([]*Message, error) {
	values := &url.Values{}
	values.Set("sid", s.secret.Sid)
	values.Set("skey", s.secret.Skey)
	values.Set("lang", "en_US")
	values.Set("pass_ticket", s.secret.PassTicket)
	url := fmt.Sprintf("%s/webwxsync?%s", s.secret.BaseUri, values.Encode())
	b, err := s.httpClient.PostJson(url, map[string]interface{}{
		"BaseRequest": s.baseRequest,
		"SyncKey":     s.secret.SyncKey,
		"rr":          ^int(time.Now().Unix()) + 1,
	})
	if err != nil {
		return nil, err
	}

	var r SyncResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.BaseResponse.Ret != 0 {
		log.Infof(string(b))
		return nil, errors.New("sync error")
	}
	s.updateSyncKey(r.SyncKey)
	return r.MsgList, nil
}
func (s *WxService) TestSyncCheck() error {
	for _, h := range []string{"webpush.", "webpush2."} {
		s.secret.PushHost = h + s.secret.Host
		syncStatus, err := s.SyncCheck()
		if err == nil {
			if syncStatus.Retcode == "0" {
				return nil
			}
		}
	}
	return errors.New("Test SyncCheck error")
}
func (s *WxService) SyncCheck() (*syncStatus, error) {
	uri := fmt.Sprintf("https://%s/cgi-bin/mmwebwx-bin/synccheck", s.secret.PushHost)
	values := &url.Values{}
	values.Set("r", TimestampStr())
	values.Set("sid", s.secret.Sid)
	values.Set("uin", strconv.FormatInt(s.secret.Uin, 10))
	values.Set("skey", s.secret.Skey)
	values.Set("deviceid", s.secret.DeviceID)
	values.Set("synckey", s.secret.SyncKeyStr)
	values.Set("_", TimestampStr())

	b, err := s.httpClient.Get(uri, values)
	if err != nil {
		return nil, err
	}
	str := string(b)
	re := regexp.MustCompile(`window.synccheck=\{retcode:"(\d+)",selector:"(\d+)"\}`)
	matchs := re.FindStringSubmatch(str)
	if len(matchs) == 0 {
		log.Infof(str)
		return nil, errors.New("find Sync check code error")
	}
	syncStatus := &syncStatus{Retcode: matchs[1], Selector: matchs[2]}
	return syncStatus, nil
}

//获取联系人
func (s *WxService) GetContacts() error {
	values := &url.Values{}
	values.Set("seq", "0")
	values.Set("pass_ticket", s.secret.PassTicket)
	values.Set("skey", s.secret.Skey)
	values.Set("r", TimestampStr())
	url := fmt.Sprintf("%s/webwxgetcontact?%s", s.secret.BaseUri, values.Encode())
	b, err := s.httpClient.PostJson(url, map[string]interface{}{})
	if err != nil {
		return err
	}
	var r ContactResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	if r.BaseResponse.Ret != 0 {
		return errors.New("Get Contacts error")
	}
	log.Infof("update %d contacts", r.MemberCount)
	s.contacts = make(map[string]*User, r.MemberCount)
	return s.updateContacts(r.MemberList)
}

//更新联系人
func (s *WxService) updateContacts(us []*User) error {
	for _, u := range us {
		s.contacts[u.UserName] = u
	}
	b, err := json.Marshal(us)
	if err != nil {
		log.Errorf("save contacts json encode error:", err)
	}
	err = ioutil.WriteFile("wx-contacts.json", b, 0644)
	if err != nil {
		log.Errorf("save json write to file error:", err)
	}
	return nil
}

//暂时用不到
func (s *WxService) StatusNotify() error {
	values := &url.Values{}
	values.Set("lang", "zh_CN")
	values.Set("pass_ticket", s.secret.PassTicket)
	url := fmt.Sprintf("%s/webwxstatusnotify?%s", s.secret.BaseUri, values.Encode())
	b, err := s.httpClient.PostJson(url, map[string]interface{}{
		"BaseRequest":  s.baseRequest,
		"code":         3,
		"FromUserName": s.user.UserName,
		"ToUserName":   s.user.UserName,
		"ClientMsgId":  TimestampMicroSecond(),
	})
	if err != nil {
		return err
	}
	return s.CheckCode(b, "Status Notify error")
}
func (s *WxService) CheckCode(b []byte, errmsg string) error {
	var r InitResponse
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	if r.BaseResponse.Ret != 0 {
		return errors.New("Status Notify error")
	}
	return nil
}
func (s *WxService) updateSyncKey(syncs *SyncKey) {
	s.secret.SyncKey = syncs
	syncKeys := make([]string, syncs.Count)
	for n, k := range syncs.List {
		syncKeys[n] = fmt.Sprintf("%d_%d", k.Key, k.Val)
	}
	s.secret.SyncKeyStr = strings.Join(syncKeys, "|")
}
func (s *WxService) NewLoginPage(newLoginUri string) error {
	b, err := s.httpClient.Get(newLoginUri+"&fun=new", nil)
	if err != nil {
		log.Infof("HTTP GET err: %s", err.Error())
		return err
	}
	err = xml.Unmarshal(b, s.secret)
	if err != nil {
		log.Infof("parse wxSecret from xml failed: %v", err)
		return err
	}
	if s.secret.Code == "0" {
		u, _ := url.Parse(newLoginUri)
		s.secret.BaseUri = newLoginUri[:strings.LastIndex(newLoginUri, "/")]
		s.secret.Host = u.Host
		s.secret.DeviceID = "e" + RandNumbers(15)
		return nil
	} else {
		return errors.New("Get wxSecret Error")
	}

}

func (s *WxService) GetNewLoginUrl() (string, error) {
	uuid, err := s.getUuid()
	if err != nil {
		return "", err
	}
	err = s.ShowQRcodeUrl(uuid)
	if err != nil {
		return "", err
	}
	newLoginUri, err := s.WaitingForLoginConfirm(uuid)
	if err != nil {
		return "", err
	}
	return newLoginUri, nil
}
func (s *WxService) WaitingForLoginConfirm(uuid string) (string, error) {
	re := regexp.MustCompile(`window.code=([0-9]*);`)
	tip := "1"
	for {
		values := &url.Values{}
		values.Set("uuid", uuid)
		values.Set("tip", tip)
		values.Set("_", TimestampStr())
		b, err := s.httpClient.Get("https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login", values)
		if err != nil {
			log.Errorf("HTTP GET err: %s", err.Error())
			return "", err
		}
		s := string(b)
		codes := re.FindStringSubmatch(s)
		if len(codes) == 0 {
			log.Errorf("find window.code failed, origin response: %s\n", s)
			return "", ErrUnknow
		} else {
			code := codes[1]
			if code == "408" {
				log.Info("login timeout, reconnecting...")
				// }else if code == "400" {
				// 	log.Info("login timeout, need refresh qrcode")
			} else if code == "201" {
				log.Info("scan success, please confirm login on your phone")
				tip = "0"
			} else if code == "200" {
				log.Info("login success")
				re := regexp.MustCompile(`window\.redirect_uri="(.*?)";`)
				us := re.FindStringSubmatch(s)
				if len(us) == 0 {
					log.Info(s)
					return "", errors.New("find redirect uri failed")
				}
				return us[1], nil
			} else {
				log.Errorf("unknow window.code %s\n", code)
				return "", ErrUnknow
			}
		}
	}
	return "", nil
}

func (s *WxService) ShowQRcodeUrl(uuid string) error {
	uri := fmt.Sprintf("%s/qrcode/%s", LoginUri, uuid)
	if s.qrcodeDir != "" {
		path, err := s.getImg(uri)
		path, _ = filepath.Abs(path)
		if err == nil {
			log.Infof("Please open img %s", path)
			return nil
		}
	}
	log.Info("Please open link in browser: " + uri)
	return nil
}
func (s *WxService) getImg(uri string) (string, error) {
	if !utils.DirExists(s.qrcodeDir) {
		err := os.MkdirAll(s.qrcodeDir, 0755)
		if err != nil {
			return "", err
		}
	}
	//https://login.weixin.qq.com/qrcode/gaWcJPzkKA==
	strs := strings.Split(uri, "/")
	name := ""
	len := len(strs)
	if len > 0 {
		name = strs[len-1]
	}
	path := fmt.Sprintf("%s/%s.jpg", s.qrcodeDir, name)
	out, err := os.Create(path)
	defer out.Close()
	b, err := s.httpClient.Get(uri, &url.Values{})
	_, err = io.Copy(out, bytes.NewReader(b))
	return path, err

}

func (s *WxService) getUuid() (string, error) {
	values := &url.Values{}
	values.Set("appid", "wx782c26e4c19acffb")
	values.Set("fun", "new")
	values.Set("lang", "zh_CN")
	values.Set("_", TimestampStr())
	uri := fmt.Sprintf("%s/jslogin", LoginUri)
	b, err := s.httpClient.Get(uri, values)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`"([\S]+)"`)
	find := re.FindStringSubmatch(string(b))
	if len(find) > 1 {
		return find[1], nil
	} else {
		return "", fmt.Errorf("get uuid error, response: %s", b)
	}
}
