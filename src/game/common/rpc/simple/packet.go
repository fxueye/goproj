package simple

import (
	"bytes"
	"encoding/binary"
)

const (
	PACKET_DEFAULT_LEN = 256
	PACKET_MAX_LEN     = 1024 * 1024
)

type Packet struct {
	buffer *bytes.Buffer
}

func NewPacket() *Packet {
	return &Packet{bytes.NewBuffer(make([]byte, 0, PACKET_DEFAULT_LEN))}
}
func NewPacketByLength(length int32) *Packet {
	if length > PACKET_MAX_LEN {
		panic("init len is larger than max length")
	}
	return &Packet{bytes.NewBuffer(make([]byte, 0, length))}
}
func NewPacketByBytes(buf []byte) *Packet {
	if len(buf) > PACKET_MAX_LEN {
		panic("byte array is larger than max length")
	}
	return &Packet{bytes.NewBuffer(buf)}
}

func (p *Packet) Len() int {
	return p.buffer.Len()
}

func (p *Packet) Bytes() []byte {
	return p.buffer.Bytes()
}

func (p *Packet) PopBool() bool {
	var v byte
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return v > 0
}
func (p *Packet) PopInt16() int16 {
	var v int16
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return v
}
func (p *Packet) PopInt32() int32 {
	var v int32
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return v
}
func (p *Packet) PopInt64() int64 {
	var v int64
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return v
}
func (p *Packet) PopFloat32() float32 {
	var v float32
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return v
}
func (p *Packet) PopFloat64() float64 {
	var v float64
	binary.Read(p.buffer, binary.LittleEndian, &v)
	return v
}
func (p *Packet) PopString() string {
	length := p.PopInt16()
	bs := make([]byte, length)
	_, err := p.buffer.Read(bs)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func (p *Packet) PutBool(v bool) {
	if v {
		binary.Write(p.buffer, binary.LittleEndian, byte(1))
	} else {
		binary.Write(p.buffer, binary.LittleEndian, byte(0))
	}
}
func (p *Packet) PutInt16(v int16) {
	binary.Write(p.buffer, binary.LittleEndian, v)
}
func (p *Packet) PutInt32(v int32) {
	binary.Write(p.buffer, binary.LittleEndian, v)
}
func (p *Packet) PutInt64(v int64) {
	binary.Write(p.buffer, binary.LittleEndian, v)
}
func (p *Packet) PutFloat32(v float32) {
	binary.Write(p.buffer, binary.LittleEndian, v)
}
func (p *Packet) PutFloat64(v float64) {
	binary.Write(p.buffer, binary.LittleEndian, v)
}
func (p *Packet) PutString(v string) {
	bs := []byte(v)
	l := int16(len(bs))
	p.PutInt16(l)
	p.buffer.Write(bs)
}
func (p *Packet) PutBytes(v []byte) {
	p.buffer.Write(v)
}
