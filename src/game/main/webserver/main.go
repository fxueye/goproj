package main

import (
	"fmt"
	"game/common/utils"
	web "game/web"
)

func main() {
	err := utils.SendMail("568669736@qq.com", "jyskjugscvgpbcde", "smtp.qq.com:587", "281431280@qq.com", "server crash", "你好", "")
	if err != nil {
		fmt.Print(err)
	}
	web.Run()
}
