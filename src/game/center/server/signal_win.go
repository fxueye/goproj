// +build windows

package server

import (
	"os"

	log "github.com/cihub/seelog"
)

func OnSignal(sig os.Signal) {
	log.Infof("%v", sig)
}
