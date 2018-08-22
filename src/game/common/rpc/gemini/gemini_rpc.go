package gemini

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
	tcp "game/common/server/tcp"
)

const (
	MAX_PACKET_SIZE = 1024 * 1024
	TCP_REQ         = 0
	TCP_ACK         = 1

	HEADER_LEN       = 1 + 4 + 4     // isAck(bool)+seq(int32)+bodyLen(int32)
	HEADER_LEN_CHECK = 1 + 4 + 4 + 4 // isAck(bool)+seq(int32)+bodyLen(int32)
)

type GeminiRPC struct {
	tcp.IProtocol
	needCheck    bool
	needHashcode bool
	timeout      time.Duration
	seq          int32
	seqMutex     sync.Mutex
	acks         map[int32]chan *GeminiResponse
	handler      *GeminiHandler
}

func NewGeminiRPC(handler *GeminiHandler, needCheck bool, timeout time.Duration) *GeminiRPC {
	rpc := new(GeminiRPC)
	rpc.needCheck = needCheck
	rpc.timeout = timeout
	rpc.handler = handler
	rpc.acks = make(map[int32]chan *GeminiResponse)
	return rpc
}

func (rpc *GeminiRPC) SetTimeout(timeout time.Duration) {
	rpc.timeout = timeout
}

func (rpc *GeminiRPC) SetNeedHashcode(value bool) {
	rpc.needHashcode = value
}

func (rpc *GeminiRPC) ResetSeq() {
	rpc.seqMutex.Lock()
	rpc.seq = 0
	rpc.seqMutex.Unlock()
}

func (rpc *GeminiRPC) ReadPack(se *tcp.Session) (tcp.IPacket, error) {
	conn := se.GetConn()
	headerLen := HEADER_LEN
	if rpc.needCheck {
		headerLen = HEADER_LEN_CHECK
	}
	headerBytes := make([]byte, headerLen)

	// read length
	if _, err := io.ReadFull(conn, headerBytes); err != nil {
		return nil, err
	}

	isAck := headerBytes[0] == TCP_ACK
	arID := int32(binary.BigEndian.Uint32(headerBytes[1:5]))
	bodyLen := int32(binary.BigEndian.Uint32(headerBytes[5:9]))

	if bodyLen > MAX_PACKET_SIZE || bodyLen < 0 {
		return nil, errors.New("the size of packet is larger than the limit")
	}
	if rpc.needCheck {
		if !checkHashCode(headerBytes[0], arID, bodyLen, headerBytes[9:]) {
			return nil, errors.New("invalid hash code")
		}
		if isAck {
			lastArID, ok := se.GetAttr("LastArID")
			if !ok {
				lastArID = int32(0)
			}
			compareArID := increSeq(lastArID.(int32))
			if arID != compareArID {
				return nil, errors.New("invalid arID on ReadPacket")
			}
			se.SetAttr("LastArID", arID)
		}
	}

	// read body
	buffBody := make([]byte, int(bodyLen))
	if _, err := io.ReadFull(conn, buffBody); err != nil {
		return nil, err
	}
	var (
		pack tcp.IPacket
		err  error
	)
	if isAck {
		pack = &GeminiRespPack{Seq: arID}
		err = pack.Decode(buffBody)
	} else {
		pack = &GeminiReqPack{Seq: arID}
		err = pack.Decode(buffBody)
	}
	if err != nil {
		return nil, err
	}
	return pack, nil
}

func (rpc *GeminiRPC) SendPack(se *tcp.Session, p tcp.IPacket) error {
	bs, err := p.Encode()
	if err != nil {
		return err
	}
	var (
		msgType   byte
		seq       int32
		sendBytes []byte
	)
	headerLen := HEADER_LEN
	if rpc.needHashcode {
		headerLen = HEADER_LEN_CHECK
	}
	sendBytes = make([]byte, headerLen+len(bs))
	// put header bytes
	if resp, ok := p.(*GeminiRespPack); ok {
		msgType = TCP_ACK
		seq = int32(resp.Seq)
	} else if req, ok := p.(*GeminiReqPack); ok {
		msgType = TCP_REQ
		seq = int32(req.Seq)
	} else {
		return errors.New("invalid packet type on SendPack")
	}
	sendBytes[0] = msgType
	binary.BigEndian.PutUint32(sendBytes[1:5], uint32(seq))
	binary.BigEndian.PutUint32(sendBytes[5:9], uint32(len(bs)))
	if rpc.needHashcode {
		codes := makeHashCode(sendBytes[0], seq, int32(len(bs)))
		copy(sendBytes[9:], codes)
	}

	// put body
	copy(sendBytes[headerLen:], bs)
	_, err = se.GetConn().Write(sendBytes)
	return err
}

func (rpc *GeminiRPC) Process(se *tcp.Session, p tcp.IPacket) error {
	if pack, ok := p.(*GeminiReqPack); ok {
		req := &GeminiRequest{
			seq:        pack.Seq,
			name:       pack.Handler,
			args:       pack.Args,
			session:    se,
			retTimeOut: rpc.timeout,
		}
		return rpc.handler.invoke(req)
	} else if pack, ok := p.(*GeminiRespPack); ok {
		if resp, ok := rpc.acks[pack.Seq]; ok {
			resp <- &GeminiResponse{
				seq:   pack.Seq,
				state: pack.State,
				err:   pack.Err,
				args:  pack.Args,
			}
			return nil
		} else {
			return errors.New(fmt.Sprintf("invalid reponse, ack=%d", pack.Seq))
		}
	}

	return errors.New("invalid packet type on Process")
}

func (rpc *GeminiRPC) Send(se *tcp.Session, handler string, args ...interface{}) error {
	pack := GeminiReqPack{
		Seq:     rpc.seq,
		Handler: handler,
		Args:    args,
	}

	return se.Send(&pack, rpc.timeout)
}

func (rpc *GeminiRPC) Transfer(uid string, se *tcp.Session, handler string, args ...interface{}) error {
	pack := GeminiReqPack{
		Seq:     rpc.seq,
		Handler: handler,
		Args:    args,
	}
	pack.Args = append(pack.Args, uid)

	return se.Send(&pack, rpc.timeout)
}

func (rpc *GeminiRPC) Call(se *tcp.Session, handler string, args ...interface{}) (*GeminiResponse, error) {
	var seq int32
	rpc.seqMutex.Lock()
	rpc.seq = increSeq(rpc.seq)
	seq = rpc.seq
	rpc.seqMutex.Unlock()
	pack := GeminiReqPack{
		Seq:     seq,
		Handler: handler,
		Args:    args,
	}
	err := se.Send(&pack, rpc.timeout)
	if err != nil {
		return nil, err
	}
	respCh := make(chan *GeminiResponse)
	rpc.acks[seq] = respCh

	select {
	case <-se.CloseChan():
		return nil, tcp.ErrConnClosing
	case resp := <-respCh:
		close(respCh)
		delete(rpc.acks, seq)
		return resp, nil
	case <-time.After(rpc.timeout):
		close(respCh)
		delete(rpc.acks, seq)
		return nil, tcp.ErrReadBlocking
	}
	return nil, tcp.ErrReadBlocking
}

func increSeq(seq int32) int32 {
	res := seq + 1
	if res >= int32(0x7FFFFFFF) {
		res = 1
	}
	return res
}

func checkHashCode(msgType byte, arID int32, bodyLen int32, compareCodes []byte) bool {
	codes := makeHashCode(msgType, arID, bodyLen)
	for i := 0; i < len(codes); i++ {
		if codes[i] != compareCodes[i] {
			return false
		}
	}
	return true
}

func makeHashCode(msgType byte, arID int32, bodyLen int32) []byte {
	src := fmt.Sprintf("%d%d%d", msgType, arID, bodyLen)
	//	hashCode := md5.Sum([]byte(src))
	h := md5.New()
	h.Write([]byte(src))
	hashCode := hex.EncodeToString(h.Sum(nil))
	return []byte(hashCode[4:8])
}

///////////////////////
// requst
type GeminiRequest struct {
	seq        int32
	session    *tcp.Session
	name       string
	args       []interface{}
	retTimeOut time.Duration
}

func (req *GeminiRequest) Name() string {
	return req.name
}
func (req *GeminiRequest) Seq() int32 {
	return req.seq
}
func (req *GeminiRequest) Args() []interface{} {
	return req.args
}
func (req *GeminiRequest) ReturnValue(args ...interface{}) error {
	pack := GeminiRespPack{
		Seq:  req.seq,
		Args: args,
	}
	return req.session.Send(&pack, req.retTimeOut)
}
func (req *GeminiRequest) ReturnError(state int32, result string) error {
	pack := GeminiRespPack{
		Seq:   req.seq,
		State: state,
		Err:   result,
	}
	return req.session.Send(&pack, req.retTimeOut)
}

func (req *GeminiRequest) ReturnResp(resp *GeminiResponse) error {
	pack := GeminiRespPack{
		Seq:  req.seq,
		Args: resp.Args(),
	}
	return req.session.Send(&pack, req.retTimeOut)
}

func (req *GeminiRequest) Session() *tcp.Session {
	return req.session
}

func (req *GeminiRequest) TimeOut() time.Duration {
	return req.retTimeOut
}

///////////////////////
// response
type GeminiResponse struct {
	seq   int32
	state int32
	err   string
	args  []interface{}
}

func (resp *GeminiResponse) Seq() int32 {
	return resp.seq
}
func (resp *GeminiResponse) State() int32 {
	return resp.state
}
func (resp *GeminiResponse) Err() string {
	return resp.err
}
func (resp *GeminiResponse) Args() []interface{} {
	return resp.args
}
