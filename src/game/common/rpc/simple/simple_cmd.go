package simple

import (
	"encoding/binary"
	tcp "game/common/server/tcp"
)

type SimpleCmd struct {
	tcp.IPacket
	seqID     int16
	opcode    int16
	sessionID int64
	pack      *Packet
}

func (cmd *SimpleCmd) SeqID() int16 {
	return cmd.seqID
}
func (cmd *SimpleCmd) Opcode() int16 {
	return cmd.opcode
}
func (cmd *SimpleCmd) Pack() *Packet {
	return cmd.pack
}
func (cmd *SimpleCmd) SID() int64 {
	return cmd.sessionID
}
func (cmd *SimpleCmd) SetSID(sid int64) {
	cmd.sessionID = sid
}

func (cmd *SimpleCmd) Decode([]byte) (err error) {
	return nil
}
func (cmd *SimpleCmd) Encode() ([]byte, error) {
	hasSID := cmd.sessionID > 0
	bodyLen := 4 // seqID(2b) + opcode(2b)
	if hasSID {
		bodyLen += 8 // sessionID(8b)
	}
	var bs []byte

	if cmd.pack != nil {
		bs = cmd.pack.Bytes()
		bodyLen += len(bs)
	}
	sendBytes := make([]byte, bodyLen)
	binary.LittleEndian.PutUint16(sendBytes[0:2], uint16(cmd.seqID))
	binary.LittleEndian.PutUint16(sendBytes[2:4], uint16(cmd.opcode))

	if hasSID {
		binary.LittleEndian.PutUint64(sendBytes[4:12], uint64(cmd.sessionID))
	}

	if bs != nil {
		if hasSID {
			copy(sendBytes[12:], bs)
		} else {
			copy(sendBytes[4:], bs)
		}
	}
	return sendBytes, nil
}
