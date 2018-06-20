package tcp

import ()

type IPacket interface {
	Decode([]byte) error
	Encode() ([]byte, error)
}

type IProtocol interface {
	ReadPack(*Session) (IPacket, error)
	SendPack(*Session, IPacket) error
}

type ISessionHandler interface {
	OnConnect(*Session) bool
	OnClose(*Session)
	OnMessage(*Session, IPacket) bool
	OnError(*Session, error)
}
