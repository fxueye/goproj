package tcp

import (
	"fmt"
	server "game/common/server"
	"net"
	"time"
)

type ClientService struct {
	server.BaseService

	host     string
	port     int
	timeout  time.Duration
	protocol server.IProtocol
	handler  server.ISessionHandler
	session  *server.Session
	seConf   server.SessionConfig
}

func NewClientService(host string, port int, timeout time.Duration, protocol server.IProtocol, handler server.ISessionHandler, seConf server.SessionConfig) *ClientService {
	s := new(ClientService)
	s.host = host
	s.port = port
	s.timeout = timeout
	s.protocol = protocol
	s.handler = handler
	s.seConf = seConf
	return s
}

func (s *ClientService) Send(p server.IPacket) error {
	if s.IsClosed() || s.session == nil {
		return server.ErrConnClosing
	}
	return s.session.Send(p, s.timeout)
}

func (s *ClientService) Session() *server.Session {
	return s.session
}

func (s *ClientService) Start() error {
	s.BaseService.Start()

	var conn *net.TCPConn
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	if s.timeout > 0 {
		c, err := net.DialTimeout("tcp", addr, s.timeout)
		if err != nil {
			return err
		}
		conn = c.(*net.TCPConn)
	} else {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
		if err != nil {
			return err
		}
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return err
		}
	}

	s.session = server.NewSession(s, conn, s.protocol, s.handler, s.seConf)
	s.session.Do()

	return nil
}

func (s *ClientService) Close() {
	defer func() {
		s.session = nil
	}()
	if s.session != nil {
		s.session.Close()
	}
	s.BaseService.Close()
}
