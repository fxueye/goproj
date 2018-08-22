package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"game/common/protocol"
)

type data struct {
	ID      string
	Session string
	Meta    string
	Content string
}

func main() {
	server := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	send(conn)
	os.Exit(0)
}
func send(conn net.Conn) {
	defer conn.Close()
	for i := 0; i < 100; i++ {
		session := GetSession()
		dt := new(data)
		dt.ID = strconv.Itoa(i)
		dt.Session = session
		dt.Content = "content"
		dt.Meta = "golang"
		words, err := json.Marshal(dt)
		checkError(err)
		conn.Write(protocol.Enpack(words))
	}
}
func GetSession() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
