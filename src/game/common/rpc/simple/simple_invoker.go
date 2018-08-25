package simple

import "game/common/server"

type SimpleInvoker interface {
	Invoke(cmd *SimpleCmd, se *server.Session) error
}
