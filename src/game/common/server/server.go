package server

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime"

	//	"sync"
	"syscall"
)

type SigCallback func(sig os.Signal)

type Server struct {
	services map[string]IService
	sigCB    SigCallback
}

func NewServer() Server {
	return Server{
		services: make(map[string]IService),
		sigCB:    nil,
	}
}

func (s *Server) RegServ(name string, serv IService) error {
	if _, ok := s.services[name]; ok {
		return errors.New(fmt.Sprintf("duplicate reigister, service name=%s", name))
	}
	s.services[name] = serv
	return nil
}

func (s *Server) GetServ(name string) IService {
	if serv, ok := s.services[name]; ok {
		return serv
	}
	return nil
}

func (s *Server) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for _, v := range s.services {
		err := v.Start()
		if err != nil {
			panic(err)
		}
	}

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig)
	//	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	for {
		select {
		case sig := <-chSig:
			{
				if s.OnSignal(sig) {
					s.Close()
					return
				}
			}
		}
	}

}

func (s *Server) RegSigCallback(cb SigCallback) {
	s.sigCB = cb
}

//返回true则关闭
func (s *Server) OnSignal(sig os.Signal) bool {
	if s.sigCB != nil {
		s.sigCB(sig)
	}

	return sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGKILL
}

func (s *Server) Close() {
	//	wg := sync.WaitGroup{}
	//	for _, v := range s.services {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			v.Close()
	//		}()
	//	}
	//	wg.Wait()

	for _, v := range s.services {
		v.Close()
	}
}
