package db

import (
	"game/common/db"
	"time"
)

type DBMgr struct {
	WxDB  *db.DBInstance
	close bool
}

func (dbm *DBMgr) GetInstance() *db.DBInstance {
	return dbm.WxDB
}
func (dbm *DBMgr) Init() error {
	dbm.close = false
	go dbm.KeepAlive()
	return nil
}
func (dbm *DBMgr) CreateWxDB(ip string, port int, user string, passwd string, dbname string, maxOpen int, maxIdle int) error {
	dbm.WxDB = new(db.DBInstance)
	return dbm.WxDB.Init(ip, port, user, passwd, dbname, maxOpen, maxIdle)
}
func (dbm *DBMgr) KeepAlive() {
	for {
		if dbm.WxDB != nil {
			dbm.WxDB.Ping()
		}
		if dbm.close {
			return
		}
		time.Sleep(time.Minute * 10)
	}
}

func (dbm *DBMgr) Close() {
	dbm.close = true
	if dbm.WxDB != nil {
		dbm.WxDB.SetClose()
	}
}
