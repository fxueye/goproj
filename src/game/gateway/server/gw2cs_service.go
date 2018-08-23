package server

import (
	"sync"
	"time"
	"game/common/server/tcp"
	rpc "game/common/rpc/simple"
	cmd "game/cmds"
	log "github.com/cihub/seelog"
)

var (
	MaxMsgTime float64 = 20
)

type GW2GSService struct{
	tcp.ISessionHandler
	*tcp.ClientService
	name string
	simpleRPC *rpc.SimpleRPC
	close bool
	msgTime time.Time
	goHeart sync.Once
}
func newGW2GSService(name string,ip string,port int) *GW2GSService{
	serv := new(GW2GSService)
	serv.name = name
	inv := cmd.NewServerGWCmdsInvoker(&GwHandlers{},GwProxyHandler)
	serv.simpleRPC = rpc.NewSimpleRPC(inv,true,time.Second * 30 ,nil)
	serv.close = false
	serv.ClientService = tcp.NewClientService(ip,port,time.Second * 30 ,serv.simpleRPC,serv,tcp.SessionConfig{4,4})
	return serv
}
func (serv *GW2GSService) Start() error{
	serv.AsyncDo(func(){
		for{
			if serv.close{
				return
			}
			if serv.Session() == nil || serv.Session().IsClosed(){
				err := serv.ClientService.Start()
				if err != nil{
					log.Error(err)
					time.Sleep(time.Second)
					continue
				}
			}
			if time.Duration(time.Now().UnixNano() - serv.msgTime.UnixNano()).Seconds() > MaxMsgTime{
				log.Error("long time not accept msg form reconnecnt!")
				serv.Session().Close()
			}
			time.Sleep(time.Second * 1)
		}
	})
	return nil
}
func (serv * GW2GSService) Close(){
	serv.close = true
}
func (serv *GW2GSService)OnConnect(se *tcp.Session) bool{
	log.Infof("[%s] connect to server, addr=%v",serv.name,se.GetConn().RemoteAddr())
	serv.msgTime = time.Now()
	// serv.simpleRPC.Send()
	serv.goHeart.Do(func(){
		go serv.CheckHeartBeat()
	})
	return true
}
func (serv *GW2GSService) CheckHeartBeat(){
	defer func(){
		if err:= recover(); err != nil{
			log.Error(err)
			ShowStack()
		}
	}()
	for {
		if serv.IsClosed(){
			time.Sleep(time.Second)
			continue
		}
		serv.simpleRPC.Send(serv.Session(),0,cmd.ServerGWCmds_HEART_BEAT,0)
		time.Sleep(time.Second * 5)
	}
}
func (serv *GW2GSService) OnClose(*tcp.Session){
	log.Infof("[%s] session closed",serv.name)
}
func (serv * GW2GSService) OnMessage(se * tcp.Session,p tcp.IPacket) bool{
	serv.msgTime = time.Now()
	defer func(){
		if err := recover(); err != nil{
			log.Error(err)
			se.Close()
		}
	}()
	serv.simpleRPC.Process(se,p)
	return true
}
func (serv *GW2GSService) Notify(sid int64, opcode int16, args ...interface{}) error {
	err := serv.simpleRPC.Send(serv.Session(), 0, opcode, sid, args...)
	if err != nil {
		log.Error(err)
		if serv.Session() != nil {
			serv.Session().Close()
		}
		return err

	}
	return nil
}
