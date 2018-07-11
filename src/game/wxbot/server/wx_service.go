package server

import (
	"game/common/server/wx"
)

type WxService struct {
	*wx.WxService
}

func newWxService(qrcodeDir string) *WxService {
	s := new(WxService)
	s.WxService = wx.NewWxService(qrcodeDir)
	return s
}
