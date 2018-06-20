package web

import (
	"game/common/web"
)

func hello(val string) string { return "hello " + val + "\n" }

func Start() {

	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
	web.Close()
}
