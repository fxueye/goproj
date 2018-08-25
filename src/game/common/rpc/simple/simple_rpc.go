package simple

import (
	"encoding/binary"
	"errors"
	"fmt"
	"game/common/server"
	utils "game/common/utils"
	"io"
	"time"
)

type SimpleRPC struct {
	server.IProtocol
	timeout       time.Duration
	invoker       SimpleInvoker
	desKey        []byte
	readSidOnRecv bool
}

func NewSimpleRPC(invoker SimpleInvoker, readSidOnRecv bool, timeout time.Duration, desKey []byte) *SimpleRPC {
	rpc := new(SimpleRPC)
	rpc.timeout = timeout
	rpc.invoker = invoker
	rpc.readSidOnRecv = readSidOnRecv
	rpc.desKey = desKey
	return rpc
}
func (rpc *SimpleRPC) SetTimeout(timeout time.Duration) {
	rpc.timeout = timeout
}

func (rpc *SimpleRPC) ReadPack(se *server.Session) (server.IPacket, error) {
	conn := se.GetConn()
	headerBytes := make([]byte, 4)

	// read length
	if _, err := io.ReadFull(conn, headerBytes); err != nil {
		return nil, err
	}
	bodyLen := int(binary.LittleEndian.Uint32(headerBytes))

	if bodyLen > PACKET_MAX_LEN || bodyLen < 0 {
		return nil, errors.New("the size of packet is larger than the limit")
	}

	// read body
	buffBody := make([]byte, int(bodyLen))
	if _, err := io.ReadFull(conn, buffBody); err != nil {
		return nil, err
	}

	if rpc.desKey != nil {
		var err error
		//		fmt.Println(len(buffBody))
		buffBody, err = utils.DesDecrypt(buffBody, rpc.desKey, nil)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		//		fmt.Println(len(buffBody))
	}
	pack := NewPacketByBytes(buffBody)

	cmd := &SimpleCmd{
		seqID:  pack.PopInt16(),
		opcode: pack.PopInt16(),
		pack:   pack,
	}
	if rpc.readSidOnRecv {
		cmd.sessionID = pack.PopInt64()
	}

	return cmd, nil
}

func (rpc *SimpleRPC) SendPack(se *server.Session, p server.IPacket) error {
	bs, err := p.Encode()
	if err != nil {
		return err
	}

	var sendBytes []byte
	if rpc.desKey != nil {
		bs, err = utils.DesEncrypt(bs, rpc.desKey, nil)
		if err != nil {
			fmt.Println(err)
			return err
		}
		sendBytes = make([]byte, len(bs)+4)
		binary.LittleEndian.PutUint32(sendBytes[:4], uint32(len(bs)))
		copy(sendBytes[4:], bs)
	} else {
		sendBytes = make([]byte, len(bs)+4)
		binary.LittleEndian.PutUint32(sendBytes[:4], uint32(len(bs)))
		copy(sendBytes[4:], bs)
	}
	_, err = se.GetConn().Write(sendBytes)
	return err
}

func (rpc *SimpleRPC) Send(se *server.Session, seqID int16, opcode int16, sid int64, args ...interface{}) error {
	if se == nil {
		return errors.New("Session Invalid")
	}
	var pack *Packet
	if args != nil && len(args) > 0 {
		pack = NewPacket()
		for i := 0; i < len(args); i++ {
			arg := args[i]
			switch t := arg.(type) {
			case bool:
				pack.PutBool(t)
			case int16:
				pack.PutInt16(t)
			case int:
				pack.PutInt32(int32(t))
			case int32:
				pack.PutInt32(int32(t))
			case uint:
				pack.PutInt32(int32(t))
			case int64:
				pack.PutInt64(int64(t))
			case uint64:
				pack.PutInt64(int64(t))
			case float32:
				pack.PutFloat32(t)
			case float64:
				pack.PutFloat64(t)
			case string:
				pack.PutString(t)
			case []byte:
				pack.PutBytes(t)
			case Wrapper:
				t.Encode(pack)
			default:
				return errors.New(fmt.Sprint("invalid argument, type=", t))
			}
		}
	}
	cmd := SimpleCmd{
		seqID:     seqID,
		opcode:    opcode,
		pack:      pack,
		sessionID: sid,
	}
	return se.Send(&cmd, rpc.timeout)
}

func (rpc *SimpleRPC) Process(se *server.Session, p server.IPacket) error {
	if cmd, ok := p.(*SimpleCmd); ok {
		return rpc.invoker.Invoke(cmd, se)
	}

	return errors.New("invalid packet type on Process")
}
