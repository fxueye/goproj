package gemini

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

var (
	intArr    = []int{1, -1, 230, -120, 30000, -30000, 80000000000, -80000000000}
	uintArr   = []uint{1, 230, 120, 30000, 80000000000}
	strArr    = []string{"aaa", "aa中文", "balalalalalalalalala巴拉拉拉拉巴拉拉拉拉巴拉拉拉拉巴拉拉拉拉"}
	real32Arr = []float32{888888888888.33, -0.888888888888}
	real64Arr = []float64{888888888888.33, -8888888888.88}
)

func TestInteger(t *testing.T) {
	for i := 0; i < len(intArr); i++ {
		str := fmt.Sprint("org: ", intArr[i])
		buff := bytes.NewBuffer(make([]byte, 0, 32))
		packb(buff, reflect.Indirect(reflect.ValueOf(intArr[i])))

		bs := buff.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		buff = bytes.NewBuffer(bs)
		v := unpackb(buff)
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestUnsignedInteger(t *testing.T) {
	for i := 0; i < len(uintArr); i++ {
		str := fmt.Sprint("org: ", uintArr[i])
		buff := bytes.NewBuffer(make([]byte, 0, 32))
		packb(buff, reflect.Indirect(reflect.ValueOf(uintArr[i])))

		bs := buff.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		buff = bytes.NewBuffer(bs)
		v := unpackb(buff)
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Log(str)
	}
}

func TestFloat32(t *testing.T) {
	for i := 0; i < len(real32Arr); i++ {
		str := fmt.Sprint("org: ", real32Arr[i])
		buff := bytes.NewBuffer(make([]byte, 0, 32))
		packb(buff, reflect.Indirect(reflect.ValueOf(real32Arr[i])))

		bs := buff.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		buff = bytes.NewBuffer(bs)
		v := unpackb(buff)
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Logf(str)
	}
}

func TestFloat64(t *testing.T) {
	for i := 0; i < len(real64Arr); i++ {
		str := fmt.Sprint("org: ", real64Arr[i])
		buff := bytes.NewBuffer(make([]byte, 0, 32))
		packb(buff, reflect.Indirect(reflect.ValueOf(real64Arr[i])))

		bs := buff.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		buff = bytes.NewBuffer(bs)
		v := unpackb(buff)
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Logf(str)
	}
}
func TestString(t *testing.T) {
	for i := 0; i < len(strArr); i++ {
		str := fmt.Sprint("org: ", strArr[i])
		buff := bytes.NewBuffer(make([]byte, 0, 32))
		packb(buff, reflect.Indirect(reflect.ValueOf(strArr[i])))

		bs := buff.Bytes()
		str = fmt.Sprint(str, fmt.Sprint(", packb: bytes=", bs))

		buff = bytes.NewBuffer(bs)
		v := unpackb(buff)
		str = fmt.Sprint(str, fmt.Sprint(", unpackb: value=", v))
		t.Logf(str)
	}
}

func TestArray(t *testing.T) {
	var arr []interface{}
	arr = []interface{}{
		intArr,
		strArr,
		real32Arr,
		real64Arr,
	}
	str := fmt.Sprint("\norg: ", arr)
	buff := bytes.NewBuffer(make([]byte, 0, 32))
	packb(buff, reflect.Indirect(reflect.ValueOf(arr)))

	bs := buff.Bytes()
	str = fmt.Sprint(str, fmt.Sprint("\n packb: bytes=", bs))

	buff = bytes.NewBuffer(bs)
	v := unpackb(buff)
	str = fmt.Sprint(str, fmt.Sprint("\n unpackb: value=", v))
	t.Logf(str)
}

func TestMap(t *testing.T) {
	var arr map[interface{}]interface{}
	arr = map[interface{}]interface{}{
		"a":  intArr,
		1:    strArr,
		1.2:  real32Arr,
		true: real64Arr,
	}
	str := fmt.Sprint("\norg: ", arr)
	buff := bytes.NewBuffer(make([]byte, 0, 32))
	packb(buff, reflect.Indirect(reflect.ValueOf(arr)))

	bs := buff.Bytes()
	str = fmt.Sprint(str, fmt.Sprint("\n packb: bytes=", bs))

	buff = bytes.NewBuffer(bs)
	v := unpackb(buff)
	str = fmt.Sprint(str, fmt.Sprint("\n unpackb: value=", v))
	t.Logf(str)
}
