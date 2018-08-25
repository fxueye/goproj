package tcp

import (
	"fmt"
	server "game/common/server"
	"net"
	"time"

	log "github.com/cihub/seelog"
)

type TcpService struct {
	server.BaseService

	port          int
	acceptTimeout time.Duration
	listener      *net.TCPListener
	protocol      server.IProtocol
	handler       server.ISessionHandler
	seConf        server.SessionConfig
}

func NewTcpService(port int, acceptTimeout time.Duration, protocol server.IProtocol, handler server.ISessionHandler, seConf server.SessionConfig) *TcpService {
	s := new(TcpService)
	s.port = port
	s.acceptTimeout = acceptTimeout
	s.protocol = protocol
	s.handler = handler
	s.seConf = seConf
	return s
}

func (s *TcpService) Start() error {
	serverAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	s.listener, err = net.ListenTCP("tcp", serverAddr)
	if err != nil {
		return err
	}
	log.Infof("listen tcp, port=%v", serverAddr)
	s.BaseService.Start()

	s.AsyncDo(func() {
		defer func() {
			recover()
			if s.listener != nil {
				s.listener.Close()
				s.listener = nil
			}
		}()
		for {
			if s.IsClosed() || s.listener == nil {
				return
			}

			s.listener.SetDeadline(time.Now().Add(s.acceptTimeout))
			conn, err := s.listener.AcceptTCP()
			if err != nil {
				continue
			}
			server.NewSession(s, conn, s.protocol, s.handler, s.seConf).Do()
		}
	})
	return nil
}
