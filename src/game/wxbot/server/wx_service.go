package server

import (
	"game/common/server"
	"game/common/wx"

	log "github.com/cihub/seelog"
)

type WxService struct {
	server.BaseService
	weixin *wx.Weixin
}

func (s *WxService) Start() error {
	s.BaseService.Start()
	newLoginUri, err := s.weixin.GetNewLoginUrl()
	if err != nil {
		return err
	}
	err = s.weixin.NewLoginPage(newLoginUri)
	if err != nil {
		return err
	}
	err = s.weixin.Init()
	if err != nil {
		return err
	}
	err = s.weixin.GetContacts()
	if err != nil {
		return err
	}
	s.AsyncDo(func() {
		defer func() {
			recover()
			if s.weixin != nil {
				log.Infof("defer wx server close!")
				s.weixin = nil
			}
		}()
		err = s.weixin.Listening()
		if err != nil {
			return
		}
	})
	return nil
}

func (s *WxService) Close() {
	log.Infof("weixin server close!")
	s.BaseService.Close()
}
func newWxService() *WxService {
	s := new(WxService)
	s.weixin = wx.NewWeixin()
	return s
}
