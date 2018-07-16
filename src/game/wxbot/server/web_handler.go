package server

import (
	"net/http"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

func (*WebHandler) Api(val string) string {
	return "hello " + val + "\n"
}
func (*WebHandler) Test(w http.ResponseWriter, r *http.Request) {
	log.Info("text")
	w.Write([]byte("hello"))
}
