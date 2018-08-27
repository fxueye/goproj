package server

import (
	"bytes"
	"encoding/binary"
	"game/common/server"
	"game/common/server/web"
	"net"
	"time"

	log "github.com/cihub/seelog"
	"golang.org/x/net/websocket"
)

type WsService struct {
	*web.WebService

	port          int
	acceptTimeout time.Duration
	listener      *net.TCPListener
	protocol      server.IProtocol
	handler       server.ISessionHandler
	seConf        server.SessionConfig
}

func newWsService(port int, acceptTimeout time.Duration) *WsService {
	serv := new(WsService)
	serv.acceptTimeout = acceptTimeout
	serv.WebService = web.NewWebService(port, serv.acceptTimeout, "")
	return serv
}
func (s *WsService) Start() error {
	s.Websocket("/", websocket.Handler(s.handler_webSocket))
	return s.WebService.Start()
}

func (s *WsService) handler_webSocket(ws *websocket.Conn) {
	for {
		var data []byte
		err := websocket.Message.Receive(ws, &data)
		if err != nil {
			log.Errorf("%v", err)
			break
		}
		buff := bytes.NewBuffer(data)
		var i int32
		binary.Read(buff, binary.BigEndian, &i)
		var b bool
		binary.Read(buff, binary.BigEndian, b)
		log.Infof("%v\n", i)
		log.Infof("%v\n", b)
		s.OnReceive(ws, data)
	}
}
func (s *WsService) OnReceive(ws *websocket.Conn, data []byte) {
	// buffer *bytes.Buffer
	err := websocket.Message.Send(ws, data)
	if err != nil {
		log.Errorf("%v", err)
	}
}
