package db

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/cihub/seelog"
)

var (
	ErrorNoRow = errors.New("Query no row")
)

type DBInstance struct {
	DB           *sql.DB
	closing      bool
	queueLock    sync.Mutex
	requestQueue *list.List
	dealQueue    *list.List
	waitGroup    sync.WaitGroup
}
type QueryRequest struct {
	query string
	args  []interface{}
}

func (dbi *DBInstance) Init(ip string, port int, user string, passwd string, dbname string, maxopen int, maxIdle int) error {
	s := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, passwd, ip, port, dbname)
	db, err := sql.Open("mysql", s)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	if maxopen > 0 {
		db.SetMaxOpenConns(maxopen)
	}
	if maxIdle > 0 {
		db.SetMaxIdleConns(maxIdle)
	}
	dbi.DB = db
	dbi.closing = false
	dbi.requestQueue = new(list.List)
	dbi.dealQueue = new(list.List)
	dbi.waitGroup.Add(1)
	log.Infof("DBInstance Init %s,%d,%d", s, maxopen, maxIdle)
	return nil
}
func (dbi *DBInstance) Ping() {
	if dbi.DB != nil {

	}
}
func (dbi *DBInstance) QueryOneRow(query string, args ...interface{}) (map[string]string, error) {
	rows, err := dbi.DB.Query(query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	found := false
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		found = true
		break
	}
	retMap := make(map[string]string)
	if found {
		for i := 0; i < len(values); i++ {
			retMap[columns[i]] = string(values[i])
		}
		return retMap, nil
	} else {
		for i := 0; i < len(columns); i++ {
			retMap[columns[i]] = ""
		}
		return retMap, ErrorNoRow
	}

}
func (dbi *DBInstance) StartTx() (*sql.Tx, error) {
	return dbi.DB.Begin()
}
func (dbi *DBInstance) Prepare(query string) (*sql.Stmt, error) {
	return dbi.DB.Prepare(query)
}
func (dbi *DBInstance) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := dbi.DB.Exec(query, args...)
	if err != nil {
		log.Error(err)
	}
	return result, err
}
func (dbi *DBInstance) DealRequestInQueue() {
	defer dbi.waitGroup.Done()
	for {
		if dbi.requestQueue.Len() > 0 {
			dbi.queueLock.Lock()
			dbi.dealQueue.PushBackList(dbi.requestQueue)
			dbi.requestQueue.Init()
			dbi.queueLock.Unlock()
		}

		if dbi.dealQueue.Len() > 0 {
			for node := dbi.dealQueue.Front(); node != nil; node = node.Next() {
				q := node.Value.(*QueryRequest)
				dbi.Exec(q.query, q.args)
			}
			dbi.dealQueue.Init()
		} else {
			time.Sleep(time.Microsecond * 10)
		}
		if dbi.closing && dbi.requestQueue.Len() == 0 {
			return
		}
	}
}
