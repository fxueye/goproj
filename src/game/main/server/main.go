package main

import (
	"fmt"
	"game/common/utils"
	"net"
	"os"

	"game/common/protocol"
)

func main() {
	// service := ":1200"
	// listener, err := net.Listen("tcp", service)
	// checkErr(err)
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		continue
	// 	}
	// 	go handleConn(conn)
	// }

	contentAll := "群:[椰晖同业交流群]:[和信芽庄岘港香香17303713401]:江湖救急~超低特价明天出发16号惠享芽庄拿起护照，说走就走直客999，不含签，同行高返！没有必消！没有现收!和信国旅香香：17303713401"
	fmt.Printf("%v", utils.RemoveDuplicatesAndEmpty(utils.GetTelNum(contentAll)))

}
func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
func handleConn(conn net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	defer conn.Close()
	go reader(readerChannel)
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "connection error:", err)
			return
		}
		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}

}
func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			fmt.Println(string(data))
		}
	}
}
