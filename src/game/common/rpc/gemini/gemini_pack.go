package gemini

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	tcp "game/common/server/tcp"
)

const (
	DefaultPackCap = 256

	NilValue          = 0xc0
	TrueValue         = 0xc3
	FalseValue        = 0xc2
	SignedInt8        = 0xd0
	UnsignedInt8      = 0xcc
	SignedInt16       = 0xd1
	UnsignedInt16     = 0xcd
	SignedInt32       = 0xd2
	UnsignedInt32     = 0xce
	SignedInt64       = 0xd3
	UnsignedInt64     = 0xcf
	Real32            = 0xca
	Real64            = 0xcb
	MinimumFixedArray = 0x90
	MaximumFixedArray = 0x9f
	Array16           = 0xdc
	Array32           = 0xdd
	MinimumFixedMap   = 0x80
	MaximumFixedMap   = 0x8f
	Map16             = 0xde
	Map32             = 0xdf
	MinimumFixedRaw   = 0xa0
	MaximumFixedRaw   = 0xbf
	Str8              = 0xd9
	Raw16             = 0xda
	Raw32             = 0xdb
	Bin8              = 0xc4
	Bin16             = 0xc5
	Bin32             = 0xc6
	Ext8              = 0xc7
	Ext16             = 0xc8
	Ext32             = 0xc9
	FixExt1           = 0xd4
	FixExt2           = 0xd5
	FixExt4           = 0xd6
	FixExt8           = 0xd7
	FixExt16          = 0xd8
)

type geminiCode byte

func (code geminiCode) String() string {
	switch code {
	case NilValue:
		return "Nil"
	case TrueValue:
		return "True"
	case FalseValue:
		return "False"
	case SignedInt8:
		return "SingnedInt8"
	case UnsignedInt8:
		return "UnsignedInt8"
	case SignedInt16:
		return "SignedInt16"
	case UnsignedInt16:
		return "UnsignedInt16"
	case SignedInt32:
		return "SignedInt32"
	case UnsignedInt32:
		return "UnsignedInt32"
	case SignedInt64:
		return "SignedInt64"
	case UnsignedInt64:
		return "UnsignedInt64"
	case Real32:
		return "Real32"
	case Real64:
		return "Real64"
	case Array16:
		return "Array16"
	case Array32:
		return "Array32"
	case Map16:
		return "Map16"
	case Map32:
		return "Map32"
	case Raw16:
		return "Raw16"
	case Raw32:
		return "Raw32"
	}

	switch code & 0xF0 {
	case 0x80:
		return "FixedMap"
	case 0x90:
		return "FixedArray"
	case 0xA0, 0xB0:
		return "FixedRaw"
	}
	return "Unknown"
}

type GeminiReqPack struct {
	tcp.IPacket
	Seq     int32
	Handler string
	Args    []interface{}
}

func (p *GeminiReqPack) Decode(bs []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	buff := bytes.NewBuffer(bs)
	data := unpackb(buff)
	if list, ok := data.([]interface{}); ok {
		p.Handler = list[0].(string)
		p.Args = list[1].([]interface{})
	} else {
		err = errors.New("invalid GeminiReqPack on Decode")
	}
	return err
}
func (p *GeminiReqPack) Encode() (bs []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	buff := bytes.NewBuffer(make([]byte, 0, DefaultPackCap))
	req := []interface{}{
		p.Handler,
		p.Args,
	}
	packb(buff, reflect.Indirect(reflect.ValueOf(req)))
	return buff.Bytes(), err
}

type GeminiRespPack struct {
	tcp.IPacket
	Seq   int32
	State int32
	Err   string
	Args  []interface{}
}

func (p *GeminiRespPack) Decode(bs []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	buff := bytes.NewBuffer(bs)
	data := unpackb(buff)
	if list, ok := data.([]interface{}); ok {
		resp := list[0].([]interface{})
		p.State = int32(resp[0].(int64))
		p.Err = string(resp[1].(string))
		if arg, t := list[1].([]interface{}); t {
			p.Args = arg

		} else {
			p.Args = make([]interface{}, 1)
			p.Args[0] = list[1]
		}
	} else {
		err = errors.New("invalid GeminiRespPack on Decode")
	}
	return err
}
func (p *GeminiRespPack) Encode() (bs []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = error(errors.New(fmt.Sprint(e)))
		}
	}()
	buff := bytes.NewBuffer(make([]byte, 0, DefaultPackCap))
	resq := []interface{}{
		[]interface{}{p.State, p.Err},
		p.Args,
	}
	packb(buff, reflect.Indirect(reflect.ValueOf(resq)))
	return buff.Bytes(), err
}

type GeminiExtPack struct {
	ExtType byte
	ExtData []byte
}

func packb(buff *bytes.Buffer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		packNil(buff)
	case reflect.Bool:
		packBool(buff, v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		packInteger(buff, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		packInteger(buff, int64(v.Uint()))
	case reflect.Float32:
		packFloat32(buff, float32(v.Float()))
	case reflect.Float64:
		packFloat64(buff, v.Float())
	case reflect.String:
		packString(buff, v.String())
	case reflect.Struct:
		if ext, ok := v.Interface().(GeminiExtPack); ok {
			packExt(buff, &ext)
		} else {
			panic(errors.New(fmt.Sprintf("[ packb ] invalid type on packb, type=%s", v.Type())))
		}

	case reflect.Slice, reflect.Array:
		if bs, ok := v.Interface().([]byte); ok {
			packBytes(buff, bs)
		} else {
			l := v.Len()
			switch {
			case l <= 15:
				buff.WriteByte(byte(MinimumFixedArray | l))
			case l <= (1<<16)-1:
				buff.WriteByte(byte(Array16))
				binary.Write(buff, binary.BigEndian, uint16(l))
			case l <= (1<<32)-1:
				buff.WriteByte(byte(Array32))
				binary.Write(buff, binary.BigEndian, uint32(l))
			default:
				panic(errors.New("[ packArray ] too huge array"))
			}
			for i := 0; i < l; i++ {
				packb(buff, reflect.Indirect(v.Index(i)))
			}
		}
	case reflect.Map:
		l := v.Len()
		switch {
		case l <= 15:
			buff.WriteByte(byte(MinimumFixedMap | l))
		case l <= (1<<16)-1:
			buff.WriteByte(byte(Array16))
			binary.Write(buff, binary.BigEndian, uint16(l))
		case l <= (1<<32)-1:
			buff.WriteByte(byte(Array32))
			binary.Write(buff, binary.BigEndian, uint32(l))
		default:
			panic(errors.New("[ packArray ] too huge array"))
		}
		keys := v.MapKeys()
		for i := 0; i < l; i++ {
			packb(buff, reflect.Indirect(keys[i]))
			packb(buff, reflect.Indirect(v.MapIndex(keys[i])))
		}
	case reflect.Interface:
		i := v.Interface()
		packb(buff, reflect.Indirect(reflect.ValueOf(i)))
	default:
		panic(errors.New(fmt.Sprintf("[ packb ] invalid type on packb, v=%v, v.Kind=%v", v, v.Kind())))
	}
}

func packNil(buff *bytes.Buffer) {
	buff.WriteByte(byte(NilValue))
}
func packBool(buff *bytes.Buffer, b bool) {
	if b {
		buff.WriteByte(byte(TrueValue))
	} else {
		buff.WriteByte(byte(FalseValue))
	}
}
func packByte(buff *bytes.Buffer, b byte) {
	buff.WriteByte(b)
}
func packInteger(buff *bytes.Buffer, i int64) {
	if i < 0 {
		switch {
		case i >= -32:
			buff.WriteByte(byte(i))
		case i >= -(1 << 7):
			buff.WriteByte(byte(SignedInt8))
			binary.Write(buff, binary.BigEndian, byte(i))
		case i >= -(1 << 15):
			buff.WriteByte(byte(SignedInt16))
			binary.Write(buff, binary.BigEndian, int16(i))
		case i >= -(1 << 31):
			buff.WriteByte(byte(SignedInt32))
			binary.Write(buff, binary.BigEndian, int32(i))
		case i >= -(1 << 63):
			buff.WriteByte(byte(SignedInt64))
			binary.Write(buff, binary.BigEndian, int64(i))
		default:
			panic(errors.New("[ packInteger ] too huge int"))
		}
	} else {
		switch {
		case i <= 127:
			buff.WriteByte(byte(i))
		case i <= (1<<8)-1:
			buff.WriteByte(byte(UnsignedInt8))
			binary.Write(buff, binary.BigEndian, byte(i))
		case i <= (1<<16)-1:
			buff.WriteByte(byte(UnsignedInt16))
			binary.Write(buff, binary.BigEndian, uint16(i))
		case i <= (1<<32)-1:
			buff.WriteByte(byte(UnsignedInt32))
			binary.Write(buff, binary.BigEndian, uint32(i))
		case uint64(i) <= uint64((1<<64)-1):
			buff.WriteByte(byte(UnsignedInt64))
			binary.Write(buff, binary.BigEndian, uint64(i))
		default:
			panic(errors.New("[ packInteger ] too huge int"))
		}
	}

}
func packFloat32(buff *bytes.Buffer, f float32) {
	buff.WriteByte(byte(Real32))
	binary.Write(buff, binary.BigEndian, f)
}
func packFloat64(buff *bytes.Buffer, d float64) {
	buff.WriteByte(byte(Real64))
	binary.Write(buff, binary.BigEndian, d)
}
func packString(buff *bytes.Buffer, str string) {
	b := []byte(str)
	l := len(b)
	switch {
	case l <= 31:
		buff.WriteByte(byte(MinimumFixedRaw | l))
		buff.Write(b)
	case l <= (1<<8)-1:
		buff.WriteByte(byte(Str8))
		buff.WriteByte(byte(l))
		buff.Write(b)
	case l <= (1<<16)-1:
		buff.WriteByte(byte(Raw16))
		binary.Write(buff, binary.BigEndian, uint16(l))
		buff.Write(b)
	case 1 <= (1<<32)-1:
		buff.WriteByte(byte(Raw32))
		binary.Write(buff, binary.BigEndian, uint32(l))
		buff.Write(b)
	default:
		panic(errors.New("[ packString ] too huge string"))
	}
}
func packBytes(buff *bytes.Buffer, bs []byte) {
	l := len(bs)
	switch {
	case l <= (1<<8)-1:
		buff.WriteByte(byte(Bin8))
		buff.WriteByte(byte(l))
		buff.Write(bs)
	case l <= (1<<16)-1:
		buff.WriteByte(byte(Bin16))
		binary.Write(buff, binary.BigEndian, uint16(l))
		buff.Write(bs)
	case 1 <= (1<<32)-1:
		buff.WriteByte(byte(Bin32))
		binary.Write(buff, binary.BigEndian, uint32(l))
		buff.Write(bs)
	default:
		panic(errors.New("[ packBytes ] too huge bytes"))
	}

}
func packOldSpecRaw(buff *bytes.Buffer, str string) {
	b := []byte(str)
	l := len(b)
	switch {
	case l <= 31:
		buff.WriteByte(byte(MinimumFixedRaw | l))
		buff.Write(b)
	case l <= (1<<16)-1:
		buff.WriteByte(byte(Raw16))
		binary.Write(buff, binary.BigEndian, uint16(l))
		buff.Write(b)
	case 1 <= (1<<32)-1:
		buff.WriteByte(byte(Raw32))
		binary.Write(buff, binary.BigEndian, uint32(l))
		buff.Write(b)
	default:
		panic(errors.New("[ packString ] too huge string"))
	}
}
func packExt(buff *bytes.Buffer, ext *GeminiExtPack) {
	l := len(ext.ExtData)
	switch {
	case l == 1:
		buff.WriteByte(byte(FixExt1))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l == 2:
		buff.WriteByte(byte(FixExt2))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l == 4:
		buff.WriteByte(byte(FixExt4))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l == 8:
		buff.WriteByte(byte(FixExt8))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l == 16:
		buff.WriteByte(byte(FixExt16))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l <= (1<<8)-1:
		buff.WriteByte(byte(Ext8))
		buff.WriteByte(byte(l))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l <= (1<<16)-1:
		buff.WriteByte(byte(Ext16))
		binary.Write(buff, binary.BigEndian, uint16(l))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	case l <= (1<<32)-1:
		buff.WriteByte(byte(Ext16))
		binary.Write(buff, binary.BigEndian, uint32(l))
		buff.WriteByte(byte(ext.ExtType & 0xFF))
		buff.Write(ext.ExtData)
	default:
		panic(errors.New("[ packExt ] too huge ext data"))
	}
}

func unpackb(buff *bytes.Buffer) interface{} {
	b, err := buff.ReadByte()
	if err != nil {
		panic(errors.New("[ unpackb ] error format"))
	}
	switch {
	case b == byte(NilValue):
		return nil
	case b == byte(TrueValue):
		return true
	case b == byte(FalseValue):
		return false
	case b == byte(0xc1):
		panic(errors.New("logic error, not reserved code"))
	case (b >= 0 && b <= 0x7f) || (b >= 0xcc && b <= 0xcf) ||
		(b >= 0xd0 && b <= 0xd3) || (b >= 0xe0 && b <= 0xff):
		return unpackInteger(buff, b)
	case b == Real32:
		return unpackFloat32(buff, b)
	case b == Real64:
		return unpackFloat64(buff, b)
	case (b >= 0xa0 && b <= 0xbf) || (b >= 0xd9 && b <= 0xdb):
		return unpackString(buff, b)
	//0xc4类型当作string处理
	case b == 0xc4:
		return string(unpackBytes(buff, b))
	case b >= 0xc5 && b <= 0xc6:
		return unpackBytes(buff, b)
	case b >= 0xd4 && b <= 0xd8:
		return unpackExt(buff, b)
	case (b >= 0x90 && b <= 0x9f) || (b == Array16) || (b == Array32):
		return unpackArray(buff, b)
	case (b >= 0x80 && b <= 0x8f) || (b == Map16) || (b == Map32):
		return unpackMap(buff, b)
	}

	return nil
}

func unpackInteger(buff *bytes.Buffer, b byte) int64 {
	switch {
	case byte(b&0xe0) == 0xe0:
		return int64(int8(b))
	case byte(b&0x080) == 0x00:
		return int64(b)
	case b == SignedInt8:
		v, _ := buff.ReadByte()
		return int64(int8(v))
	case b == UnsignedInt8:
		v, _ := buff.ReadByte()
		return int64(v)
	case b == SignedInt16:
		var v int16
		binary.Read(buff, binary.BigEndian, &v)
		return int64(v)
	case b == UnsignedInt16:
		var v uint16
		binary.Read(buff, binary.BigEndian, &v)
		return int64(v)
	case b == SignedInt32:
		var v int32
		binary.Read(buff, binary.BigEndian, &v)
		return int64(v)
	case b == UnsignedInt32:
		var v uint32
		binary.Read(buff, binary.BigEndian, &v)
		return int64(v)
	case b == SignedInt64 || b == UnsignedInt64:
		var v int64
		binary.Read(buff, binary.BigEndian, &v)
		return v
	}
	panic(errors.New("[ unpackInteger ] logic error, not int"))
}
func unpackString(buff *bytes.Buffer, b byte) string {
	var length int
	var bs []byte
	var err error
	switch {
	case (b & 0xe0) == 0xa0:
		length = int(b & (^byte(0xe0)))
	case b == byte(Str8):
		var l byte
		l, err = buff.ReadByte()
		length = int(l)
	case b == byte(Raw16):
		var l uint16
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	case b == byte(Raw32):
		var l uint32
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	default:
		panic(errors.New("[ unpackString ]logic error"))
	}
	if err != nil {
		panic(err)
	}
	bs = make([]byte, length)
	_, err = buff.Read(bs)
	if err != nil {
		panic(err)
	}
	return string(bs)
}
func unpackBytes(buff *bytes.Buffer, b byte) []byte {
	var length int
	var bs []byte
	var err error
	switch b {
	case byte(Bin8):
		var l byte
		l, err = buff.ReadByte()
		length = int(l)
	case byte(Bin16):
		var l uint16
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	case byte(Bin32):
		var l uint32
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	default:
		panic(errors.New("[ unpackBytes ]logic error"))
	}
	if err != nil {
		panic(err)
	}
	bs = make([]byte, length)
	_, err = buff.Read(bs)
	if err != nil {
		panic(err)
	}
	return bs
}
func unpackFloat32(buff *bytes.Buffer, b byte) float32 {
	var v float32
	err := binary.Read(buff, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}
	return v
}
func unpackFloat64(buff *bytes.Buffer, b byte) float64 {
	var v float64
	err := binary.Read(buff, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}
	return v
}
func unpackExt(buff *bytes.Buffer, b byte) *GeminiExtPack {
	var length int
	var err error
	switch b {
	case byte(FixExt1):
		length = 1
	case byte(FixExt2):
		length = 2
	case byte(FixExt4):
		length = 4
	case byte(FixExt8):
		length = 8
	case byte(FixExt16):
		length = 16
	case byte(Ext8):
		var l byte
		l, err = buff.ReadByte()
		length = int(l)
	case byte(Ext16):
		var l uint16
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	case byte(Ext32):
		var l uint32
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	default:
		panic(errors.New("[ unpackExt ]logic error"))
	}
	if err != nil {
		panic(err)
	}
	v := new(GeminiExtPack)
	v.ExtType, err = buff.ReadByte()
	if err != nil {
		panic(err)
	}
	v.ExtData = make([]byte, length)
	_, err = buff.Read(v.ExtData)
	if err != nil {
		panic(err)
	}
	return v
}

func unpackArray(buff *bytes.Buffer, b byte) []interface{} {
	var length int
	var err error
	switch {
	case (b & byte(0xf0)) == byte(0x90):
		length = int(b & (^byte(0xf0)))
	case b == byte(Array16):
		var l uint16
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	case b == byte(Array32):
		var l uint32
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	default:
		panic(errors.New("[ unpackArray ]logic error"))
	}
	if err != nil {
		panic(err)
	}
	v := make([]interface{}, length)
	for i := 0; i < length; i++ {
		v[i] = unpackb(buff)
	}
	return v
}

func unpackMap(buff *bytes.Buffer, b byte) map[interface{}]interface{} {
	var length int
	var err error
	switch {
	case (b & byte(0xf0)) == byte(0x80):
		length = int(b & (^byte(0xf0)))
	case b == byte(Map16):
		var l uint16
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	case b == byte(Map32):
		var l uint32
		err = binary.Read(buff, binary.BigEndian, &l)
		length = int(l)
	default:
		panic(errors.New("[ unpackArray ]logic error"))
	}
	if err != nil {
		panic(err)
	}
	v := make(map[interface{}]interface{})
	for i := 0; i < length; i++ {
		mk := unpackb(buff)
		mv := unpackb(buff)
		v[mk] = mv
	}
	return v
}

func try(fun func(), handler func(interface{})) {
	defer func() error {
		if err := recover(); err != nil {
			return error(errors.New(fmt.Sprint(err)))
		}
		return nil
	}()
	fun()
}
