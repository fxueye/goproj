package server

import (
	"game/common/server/web"
	"net/http"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

// func (*WebHandler) Api(val string) string {
// 	log.Infof("ctx : %v", ctx)
// 	return "hello " + val + "\n"
// }
// func (*WebHandler) Api(ctx *web.Context, val string) string {
// 	log.Infof("ctx : %v", ctx)
// 	return "hello " + val + "\n"
// }
func (*WebHandler) Api(ctx *web.Context) {

	ctx.Write([]byte(""))
}
func (*WebHandler) Test(w http.ResponseWriter, r *http.Request) {
	log.Info("text")
	w.Write([]byte("hello"))
}
