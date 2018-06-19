package handler

import (
	"chess/common"
	"chess/game/server"
)

func HandleEcho(userid uint32, connid uint32, msgBody []byte) {
	server.SendResp(userid, MsgidEchoResp, common.ResultSuccess, msgBody)
}
