package main

import (

	"fmt"
	"os"
	"time"

	"math/rand"
	"game/common/utils"
)


var (
	sendSize = 0
	isInit = false
)
func main() {
	for{
		if !isInit{
			for i := 0; i < 100000; i++{
				Send(i)
				
			}
			isInit = true
		}else{
			for i := 0; i < 1000;i++{
				Send(rand.Intn(100000-1)+1)
			}
		}
		time.Sleep(time.Second)
	}
	
}
func Send(i int){
	params := make(map[string]interface{})
	params["sessionKey"] = GetSession()
	params["openId"] = fmt.Sprintf("openid_%v",i)
	
	fmt.Printf("params:%v\n",params)
	str,_ := utils.HttpGet("http://192.168.1.188:8080/session/",params)
	fmt.Printf("ret:%v\n",str)
}


func GetSession() string {
	return fmt.Sprintf("%v%v",time.Now().Unix(),rand.Intn(99999) +100000 )
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
