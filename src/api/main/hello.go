package main

import (
	"common/web"
)

func hello(val string) string {
	return "hello"
}
func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:8081")
}
