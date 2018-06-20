package server

import (
	"game/common/server/web"
	"time"
)

type WebService struct {
	*web.WebService
}

func newWebService(port int, staticDir string) *WebService {
	serv := new(WebService)
	serv.WebService = web.NewWebService(port, time.Second, staticDir)
	return serv
}
