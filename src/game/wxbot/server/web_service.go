package server

import (
	"fmt"
	"game/common/server/web"
	"net/http"
	"reflect"
	"strings"
	"time"

	log "github.com/cihub/seelog"
)

type WebService struct {
	*web.WebService
}

func newWebService(port int, staticDir string) *WebService {
	serv := new(WebService)
	serv.WebService = web.NewWebService(port, time.Second, staticDir)
	return serv
}
func (s *WebService) Start() error {
	s.regHandlers(&WebHandler{})
	return s.WebService.Start()
}
func (s *WebService) regHandlers(handler interface{}) error {
	t := reflect.TypeOf(handler)
	v := reflect.ValueOf(handler)
	for i := 0; i < t.NumMethod(); i++ {
		if mt := t.Method(i); mt.PkgPath == "" {
			vt := v.Method(i)
			vi := vt.Interface()
			if f, ok := vi.(func(string) string); ok {
				name := fmt.Sprintf("/%s/(.*)", strings.ToLower(mt.Name))
				s.HandleFunc(name, "", f)
				continue
			}
			if f, ok := vi.(func(http.ResponseWriter, *http.Request)); ok {
				name := fmt.Sprintf("/%s", strings.ToLower(mt.Name))
				log.Infof("name:%v", name)
				s.HandleFunc(name, "", f)
				continue
			}
		}
	}
	return nil
}
