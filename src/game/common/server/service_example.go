package server

import (
	log "github.com/cihub/seelog"
	"os"
)

type ExampleService struct {
	BaseService
	count int
}

func NewExampleService() *ExampleService {
	s := new(ExampleService)
	return s
}

func (s *ExampleService) Start() error {
	s.BaseService.Start()
	log.Info("on start")
	s.count = 0
	s.AsyncDo(s.update)
	return nil
}

func (s *ExampleService) update() {
	for {
		s.count++
		log.Debugf("on update, count=%d", s.count)
		if s.count > 10000 {
			os.Exit(0)
		}
	}
}

func (s *ExampleService) Close() {
	log.Error("on close")
	s.BaseService.Close()
}
