package simple

import (
	//"bytes"
	"fmt"
	//"reflect"
	"testing"
)

var (
	boolArr   = []bool{true, false}
	int16Arr  = []int16{1, -1, 230, -120, 30000, -30000}
	intArr    = []int{1, -1, 230, -120, 30000, -30000, 100000000, -100000000}
	int64Arr  = []int64{1, -1, 230, -120, 30000, -30000, 1000000000000, -1000000000000}
	strArr    = []string{"aaa", "aa中文", "balalalalalalalalala巴拉拉拉拉巴拉拉拉拉巴拉拉拉拉巴拉拉拉拉"}
	real32Arr = []float32{1, 9, 0, 888888888888.33, -0.888888888888}
	real64Arr = []float64{1, 9, 0, 888888888888.33, -8888888888.88}
)

type TestWrap struct {
	Wrapper
	str     string
	boolArr []bool
	intArr  []int
	strArr  []string
}

func (t *TestWrap) Decode(p *Packet) {
	t.str = p.PopString()

	boolArrLen := int(p.PopInt16())
	t.boolArr = make([]bool, boolArrLen)
	for i := 0; i < boolArrLen; i++ {
		t.boolArr[i] = p.PopBool()
	}

	intArrLen := int(p.PopInt16())
	t.intArr = make([]int, intArrLen)
	for i := 0; i < intArrLen; i++ {
		t.intArr[i] = int(p.PopInt32())
	}

	strArrLen := int(p.PopInt16())
	t.strArr = make([]string, strArrLen)
	for i := 0; i < strArrLen; i++ {
		t.strArr[i] = p.PopString()
	}
}

func (t *TestWrap) Encode(p *Packet) {
	p.PutString(t.str)

	boolArrLen := len(t.boolArr)
	p.PutInt16(int16(boolArrLen))
	for i := 0; i < boolArrLen; i++ {
		p.PutBool(t.boolArr[i])
	}

	intArrLen := len(t.intArr)
	p.PutInt16(int16(intArrLen))
	for i := 0; i < intArrLen; i++ {
		p.PutInt32(int32(t.intArr[i]))
	}

	strArrLen := len(t.strArr)
	p.PutInt16(int16(strArrLen))
	for i := 0; i < strArrLen; i++ {
		p.PutString(t.strArr[i])
	}
}

func (t *TestWrap) String() string {
	return fmt.Sprint("str=", t.str, "\nboolArr=", t.boolArr, "\nintArr=", t.intArr, "\nstrArr=", t.strArr)
}

func TestBool(t *testing.T) {
	for i := 0; i < len(boolArr); i++ {
		str := fmt.Sprint("org: ", boolArr[i])
		pack := NewPacket()
		pack.PutBool(boolArr[i])

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopBool()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestShort(t *testing.T) {
	for i := 0; i < len(int16Arr); i++ {
		str := fmt.Sprint("org: ", int16Arr[i])
		pack := NewPacket()
		pack.PutInt16(int16Arr[i])

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopInt16()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestInteger(t *testing.T) {
	for i := 0; i < len(intArr); i++ {
		str := fmt.Sprint("org: ", intArr[i])
		pack := NewPacket()
		pack.PutInt32(int32(intArr[i]))

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopInt32()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < len(int64Arr); i++ {
		str := fmt.Sprint("org: ", int64Arr[i])
		pack := NewPacket()
		pack.PutInt64(int64Arr[i])

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopInt64()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestFloat32(t *testing.T) {
	for i := 0; i < len(real32Arr); i++ {
		str := fmt.Sprint("org: ", real32Arr[i])
		pack := NewPacket()
		pack.PutFloat32(real32Arr[i])

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopFloat32()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestFloat64(t *testing.T) {
	for i := 0; i < len(real64Arr); i++ {
		str := fmt.Sprint("org: ", real64Arr[i])
		pack := NewPacket()
		pack.PutFloat64(real64Arr[i])

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopFloat64()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestString(t *testing.T) {
	for i := 0; i < len(strArr); i++ {
		str := fmt.Sprint("org: ", strArr[i])
		pack := NewPacket()
		pack.PutString(strArr[i])

		bs := pack.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		pp := NewPacketByBytes(bs)
		v := pp.PopString()
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestWrapper(t *testing.T) {
	tw := &TestWrap{
		str:     "test",
		boolArr: boolArr,
		intArr:  intArr,
		strArr:  strArr,
	}

	str := fmt.Sprint("org: \n", tw)
	pack := NewPacket()
	tw.Encode(pack)

	bs := pack.Bytes()
	str = fmt.Sprint(str, fmt.Sprint("\npackb: bytes=", bs))

	pp := NewPacketByBytes(bs)
	tt := &TestWrap{}
	tt.Decode(pp)
	str = fmt.Sprint(str, fmt.Sprint("\nunpackb: value=\n", tt))
	t.Log(str)
}
