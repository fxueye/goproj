package db

import (
	//	"database/sql"
	//	"fmt"
	//	log "github.com/cihub/seelog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBMgr struct {
	UserMapDB *DBInstance
	UserDBs   map[int]*DBInstance
	close     bool
}

func (dbm *DBMgr) Init() error {
	dbm.UserDBs = make(map[int]*DBInstance)
	dbm.close = false

	go dbm.KeepAlive()
	return nil
}

func (dbm *DBMgr) CreateUserMapDB(ip string, port int, user string, passwd string, dbname string, maxOpen int, maxIdle int) error {
	dbm.UserMapDB = new(DBInstance)
	return dbm.UserMapDB.Init(ip, port, user, passwd, dbname, maxOpen, maxIdle)
}

func (dbm *DBMgr) CreateUserDB(idx int, ip string, port int, user string, passwd string, dbname string, maxOpen int, maxIdle int) error {
	dbm.UserDBs[idx] = new(DBInstance)
	return dbm.UserDBs[idx].Init(ip, port, user, passwd, dbname, maxOpen, maxIdle)
}

func (dbm *DBMgr) GetUserMapDBInstance() *DBInstance {
	return dbm.UserMapDB
}

func (dbm *DBMgr) GetUserDBInstance(idx int) *DBInstance {
	return dbm.UserDBs[idx]
}

func (dbm *DBMgr) KeepAlive() {
	for {
		if dbm.UserMapDB != nil {
			dbm.UserMapDB.Ping()
		}
		for _, v := range dbm.UserDBs {
			v.Ping()
		}
		if dbm.close {
			return
		}
		time.Sleep(time.Minute * 10)
	}
}

func (dbm *DBMgr) Close() {
	dbm.close = true
	if dbm.UserMapDB != nil {
		dbm.UserMapDB.SetClose()
	}
	for _, v := range dbm.UserDBs {
		v.SetClose()
	}
}
