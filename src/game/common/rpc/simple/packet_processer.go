package simple

import (
	//	"container/list"
	"fmt"
	//	"sync"
	tcp "game/common/server/tcp"
)

type PacketData struct {
	se *tcp.Session
	p  tcp.IPacket
}

//用于将rpc收到的packet集中处理，避免处理逻辑时考虑多线程数据访问的问题
type PacketProcesser struct {
	invoker SimpleInvoker
	//	packets  list.List
	//	lock     sync.Mutex
	packChan chan *PacketData
}

func (pp *PacketProcesser) Init(invoker SimpleInvoker, cnt int) {
	pp.invoker = invoker
	pp.packChan = make(chan *PacketData, cnt)
}

func (pp *PacketProcesser) AddPacket(se *tcp.Session, p tcp.IPacket) {
	data := &PacketData{
		se: se,
		p:  p,
	}
	//	pp.lock.Lock()
	//	defer pp.lock.Unlock()
	//	pp.packets.PushBack(data)
	pp.packChan <- data
}

func (pp *PacketProcesser) GetPacket() *PacketData {
	select {
	case v := <-pp.packChan:
		return v
	default:
		return nil
	}
}

func (pp *PacketProcesser) ProcessPacket(max int) int {
	//
	//	pp.lock.Lock()
	//	defer pp.lock.Unlock()

	cnt := 0

	//	data := pp.packets.Front()
	d := pp.GetPacket()
	for {
		if d == nil {
			return cnt
		}
		//		d := data.Value.(*PacketData)
		if cmd, ok := d.p.(*SimpleCmd); ok {
			pp.Invoke(cmd, d.se)
		}
		//		next := data.Next()
		//		pp.packets.Remove(data)
		//		data = next
		cnt++
		if cnt >= max {
			return cnt
		}
		d = pp.GetPacket()
	}
}

func (pp *PacketProcesser) Invoke(cmd *SimpleCmd, se *tcp.Session) {
	if err := recover(); err != nil {
		fmt.Println(err)
		se.Close()
	}

	pp.invoker.Invoke(cmd, se)
}
