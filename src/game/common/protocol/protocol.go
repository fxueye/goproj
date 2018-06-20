package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader       = "headers"
	ConstHeaderLength = 7
	ConstMlength      = 4
)

func Enpack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}
func Depack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)
	var i int
	data := make([]byte, 32)
	for i := 0; i < length; i++ {
		if length < i+ConstHeaderLength+ConstMlength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMlength])
			if length < i+ConstHeaderLength+ConstMlength+messageLength {
				break
			}
			data = buffer[i+ConstHeaderLength+ConstMlength : i+ConstHeaderLength+ConstMlength+messageLength]
			readerChannel <- data
			i += ConstHeaderLength + ConstMlength + messageLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}
