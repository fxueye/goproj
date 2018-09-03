package main

import (

	"fmt"
	"os"
	"strconv"
	"time"

	"game/common/utils"
)



func main() {
	for{
		for i := 0; i < 1000; i++{
			params := make(map[string]interface{})
			params["sessionKey"] = GetSession()
			params["openId"] = i
			str,_ := utils.HttpGet("http://192.168.1.188:8080/Session",params)
			fmt.Printf("ret:%v",str)
		}
		time.Sleep(time.Microsecond * 200)
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
