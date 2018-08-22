package simple

import (
	tcp "game/common/server/tcp"
)

type SimpleInvoker interface {
	Invoke(cmd *SimpleCmd, se *tcp.Session) error
}
