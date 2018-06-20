// +build !windows

package server

import (
	"os"
	"syscall"
)

func GWOnSignal(sig os.Signal) {
	switch sig {
	case syscall.SIGUSR1:
		{

		}
	default:
	}
}
