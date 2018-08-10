package simple

import ()

type Wrapper interface {
	Decode(*Packet) Wrapper
	Encode(*Packet)
}
