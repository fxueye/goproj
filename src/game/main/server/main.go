package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

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

	contentAll := "@ed5c9309bc41399af4d4b5239125d69f6d41f87aca35b25da10c42473e209eeb:<br/><span class=\"emoji emoji1f4a7\"></span>暑期大精彩 ·芽庄金兰湾<span class=\"emoji emoji1f4a7\"></span><br/><span class=\"emoji emoji1f3b5\"></span>超高体验度--传奇金兰湾<span class=\"emoji emoji1f3b5\"></span><br/><span class=\"emoji emoji26a0\"></span>无自费，❹个店<br/><span class=\"emoji emoji1f420\"></span>深度出海升级蚕岛游<br/><span class=\"emoji emoji2747\"></span>住宿：全程精选三星酒店<br/><span class=\"emoji emoji1f4e6\"></span>赠送：唯美天堂湾+精致越式SPA<br/><span class=\"emoji emoji1f4e6\"></span>赠送：佳琳娜秀+慢生活下午茶<br/>7月19号 5晚6天 同行2180<br/>7月21/26/28号 5晚6天 同行2380<br/><span class=\"emoji emoji1f449\"></span>PS:可+600全程升级四星酒店，+1000全程升级五星酒店<br/><span class=\"emoji emoji1f446\"></span>以上报价均为同行价，不含签小380结算<br/><span class=\"emoji emoji1f42c\"></span>芽庄包机中心：俊亚18137112720=微信"
	index := strings.Index(contentAll, ":")
	sendUserName := string([]byte(contentAll)[0:index])
	content := string([]byte(contentAll)[index:])
	fmt.Printf("content:%s", content)
	content = strings.Replace(content, "<br/>", "\n", -1)
	fmt.Printf("content:%s", content)
	exp := regexp.MustCompile(`<span class=".*?"></span>`)
	fmt.Printf("content:%s", content)
	content = exp.ReplaceAllString(content, "")
	fmt.Printf(sendUserName)

	fmt.Printf("content:%s", content)

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
