package server

import (
	"crypto/tls"
	"fmt"
	"game/common/server/web"
	"reflect"
	"strings"
	"time"
)

type WebTlsService struct {
	*web.WebService
}

func newWebTlsService(port int, staticDir string, cert []byte, pkey []byte) *WebTlsService {
	serv := new(WebTlsService)
	tlsConfig, _ := CreateTlsConfig(cert, pkey)
	serv.WebService = web.NewWebService(port, time.Second, staticDir, tlsConfig)
	return serv
}
func CreateTlsConfig(cert []byte, pkey []byte) (tlsConfig *tls.Config, err error) {
	config := tls.Config{
		Time: nil,
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.X509KeyPair(cert, pkey)
	if err != nil {
		println(err.Error())
		return
	}
	tlsConfig = &config
	return
}
func (s *WebTlsService) Start() error {
	s.regHandlers(&WebHandler{})
	return s.WebService.Start()
}
func (s *WebTlsService) regHandlers(handler interface{}) error {
	t := reflect.TypeOf(handler)
	v := reflect.ValueOf(handler)
	for i := 0; i < t.NumMethod(); i++ {
		if mt := t.Method(i); mt.PkgPath == "" {
			vt := v.Method(i)
			vi := vt.Interface()
			if f, ok := vi.(func(*web.Context, string)); ok {
				name := fmt.Sprintf("/%s/(.*)", strings.ToLower(mt.Name))
				s.HandleFunc(name, "", f)
				continue
			}
			if f, ok := vi.(func(*web.Context, string) string); ok {
				name := fmt.Sprintf("/%s/(.*)", strings.ToLower(mt.Name))
				s.HandleFunc(name, "", f)
				continue
			}
			if f, ok := vi.(func(string) string); ok {
				name := fmt.Sprintf("/%s/(.*)", strings.ToLower(mt.Name))
				s.HandleFunc(name, "", f)
				continue
			}
			if f, ok := vi.(func(*web.Context)); ok {
				name := fmt.Sprintf("/%s", strings.ToLower(mt.Name))
				s.HandleFunc(name, "", f)
				continue
			}

		}
	}
	return nil
}
