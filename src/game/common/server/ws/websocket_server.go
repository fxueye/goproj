package ws

import (
	"crypto/tls"
	// "bytes"
	// "encoding/binary"
	"game/common/server"
	"game/common/server/web"
	"net"
	"time"

	log "github.com/cihub/seelog"
	"golang.org/x/net/websocket"
)

type WebsocketService struct {
	*web.WebService

	port          int
	acceptTimeout time.Duration
	listener      *net.TCPListener
	protocol      server.IProtocol
	handler       server.ISessionHandler
	seConf        server.SessionConfig
}

func NewWebsocketService(port int, acceptTimeout time.Duration, protocol server.IProtocol, handler server.ISessionHandler, seConf server.SessionConfig, tlsConfig *tls.Config) *WebsocketService {
	s := new(WebsocketService)
	s.acceptTimeout = acceptTimeout
	s.WebService = web.NewWebService(port, s.acceptTimeout, "", tlsConfig)
	s.protocol = protocol
	s.handler = handler
	s.seConf = seConf
	return s
}
func (s *WebsocketService) Start() error {
	s.Websocket("/", websocket.Handler(s.handler_webSocket))
	return s.WebService.Start()
}

func (s *WebsocketService) handler_webSocket(ws *websocket.Conn) {
	log.Infof("new ws conn ip :%v", ws.RemoteAddr())

	se := server.NewWsSesion(s, ws, s.protocol, s.handler, s.seConf)
	se.Do()
	for !se.IsClosed() {
	}
	// for {
	// 	var data []byte
	// 	err := websocket.Message.Receive(ws, &data)
	// 	if err != nil {
	// 		log.Errorf("%v", err)
	// 		break
	// 	}
	// 	buff := bytes.NewBuffer(data)
	// 	var i int32
	// 	binary.Read(buff, binary.BigEndian, &i)
	// 	var b bool
	// 	binary.Read(buff, binary.BigEndian, b)
	// 	log.Infof("%v\n", i)
	// 	log.Infof("%v\n", b)
	// 	s.OnReceive(ws, data)
	// }
}
