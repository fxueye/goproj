// +build windows

package server

import (
	"os"

	log "github.com/cihub/seelog"
)

func GWOnSignal(sig os.Signal) {
	log.Infof("%v", sig)
}
