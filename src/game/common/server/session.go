package server

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/cihub/seelog"
	"golang.org/x/net/websocket"
)

// Error type
var (
	ErrConnClosing   = errors.New("use of closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking  = errors.New("read packet was blocking")
)

type SessionConfig struct {
	SendChanLimit int
	RecvChanLimit int
}

type Session struct {
	Sid    int64
	IsWs   bool
	conn   *net.TCPConn
	wsconn *websocket.Conn

	mu    sync.RWMutex //protect attrs
	attrs map[string]interface{}

	serv     IService
	protocol IProtocol
	handler  ISessionHandler

	closeOnce sync.Once     // close the conn, once, per instance
	closeFlag int32         // close flag
	closeChan chan struct{} // close chanel
	sendChan  chan IPacket  // packet send chanel
	recvChan  chan IPacket  // packet receive chanel
}

func NewSession(serv IService, conn *net.TCPConn, protocol IProtocol, handler ISessionHandler, seConf SessionConfig) *Session {
	se := new(Session)
	se.IsWs = false
	se.conn = conn
	se.attrs = make(map[string]interface{})
	se.serv = serv
	se.protocol = protocol
	se.handler = handler
	se.closeChan = make(chan struct{})
	se.sendChan = make(chan IPacket, seConf.SendChanLimit)
	se.recvChan = make(chan IPacket, seConf.RecvChanLimit)
	return se
}
func NewWsSesion(serv IService, wsconn *websocket.Conn, protocol IProtocol, handler ISessionHandler, seConf SessionConfig) *Session {
	se := new(Session)
	se.IsWs = true
	se.wsconn = wsconn
	se.attrs = make(map[string]interface{})
	se.serv = serv
	se.protocol = protocol
	se.handler = handler
	se.closeChan = make(chan struct{})
	se.sendChan = make(chan IPacket, seConf.SendChanLimit)
	se.recvChan = make(chan IPacket, seConf.RecvChanLimit)
	return se
}

func (se *Session) CloseChan() <-chan struct{} {
	return se.closeChan
}

func (se *Session) GetConn() interface{} {
	if se.IsWs {
		return se.wsconn
	} else {
		return se.conn
	}
}

func (se *Session) SetAttr(name string, value interface{}) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.attrs[name] = value
}

func (se *Session) GetAttr(name string) (interface{}, bool) {
	se.mu.RLock()
	defer se.mu.RUnlock()
	v, ok := se.attrs[name]
	return v, ok
}

// IsClosed indicates whether or not the connection is closed
func (se *Session) IsClosed() bool {
	return atomic.LoadInt32(&se.closeFlag) == 1
}

// Close closes the connection
func (se *Session) Close() {
	if se.IsClosed() {
		return
	}

	se.closeOnce.Do(func() {
		atomic.StoreInt32(&se.closeFlag, 1)
		close(se.closeChan)
		close(se.sendChan)
		close(se.recvChan)
		if se.IsWs {
			se.wsconn.Close()
		} else {
			se.conn.Close()
		}
		se.handler.OnClose(se)
	})
}

func (se *Session) Do() {
	if !se.handler.OnConnect(se) {
		return
	}
	se.serv.AsyncDo(se.messageLoop)
	se.serv.AsyncDo(se.readLoop)
	se.serv.AsyncDo(se.writeLoop)
}

func (se *Session) Send(p IPacket, timeout time.Duration) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = ErrConnClosing
		}
	}()
	if se.IsClosed() {
		return ErrConnClosing
	}
	if timeout <= 0 {
		select {
		case se.sendChan <- p:
			return nil
		default:
			return ErrWriteBlocking
		}
	} else {
		select {
		case se.sendChan <- p:
			return nil
		case <-se.closeChan:
			return ErrConnClosing
		case <-time.After(timeout):
			return ErrWriteBlocking
		}
	}
}

func (se *Session) fix() {
	recover()
	se.Close()
}

func (se *Session) readLoop() {

	defer se.fix()
	for {
		if se.serv.IsClosed() {
			return
		}
		select {
		case <-se.closeChan:
			return
		default:
		}

		p, err := se.protocol.ReadPack(se)
		if err != nil {
			log.Error(err)
			return
		}

		se.recvChan <- p
	}
}
func (se *Session) writeLoop() {
	defer se.fix()
	for {
		if se.serv.IsClosed() {
			return
		}
		select {
		case <-se.closeChan:
			return
		case p := <-se.sendChan:
			if se.IsClosed() {
				return
			}
			if err := se.protocol.SendPack(se, p); err != nil {
				log.Error(err)
				return
			}
		}
	}
}
func (se *Session) messageLoop() {
	defer se.fix()

	for {
		if se.serv.IsClosed() {
			return
		}
		select {
		case <-se.closeChan:
			return
		case p := <-se.recvChan:
			if se.IsClosed() {
				return
			}
			if !se.handler.OnMessage(se, p) {
				return
			}
		}
	}
}
