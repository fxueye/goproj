package wx

import (
	"errors"
	"game/common/server"
)

var (
	fs                   = &FileStore{"/tmp/wx.secret"}
	LoginUri             = "https://login.weixin.qq.com"
	ErrUnknow            = errors.New("Unknow Error")
	ErrUserNotExists     = errors.New("Error User Not Exist")
	ErrNotLogin          = errors.New("Not Login")
	ErrLoginTimeout      = errors.New("Login Timeout")
	ErrWaitingForConfirm = errors.New("Waiting For Confirm")
)

type WxService struct {
	server.BaseService

	httpClient  *Client
	secret      *wxSecret
	baseRequest *BaseRequest
	user        *User
	contacts    map[string]*User
}

func NewWxService() *WxService {
	s := new(WxService)
	s.httpClient = NewClient()
	s.secret = &wxSecret{}
	s.baseRequest = &BaseRequest{}
	s.user = &User{}
	s.contacts = make(map[string]*User)
	return s
}
