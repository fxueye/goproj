package db

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
)

const (
	Statement_Insert = 1
	Statement_Load   = 2
)

var (
	ErrorNoRow = errors.New("Query no row")
)

//type TableStatementGroup struct {
//	Stmts map[int]*sql.Stmt
//}
//
type QueryRequest struct {
	query string
	args  []interface{}
}
type DBInstance struct {
	MyDB         *sql.DB
	closing      bool
	queueLock    sync.Mutex
	requestQueue *list.List
	dealQueue    *list.List
	waitGroup    sync.WaitGroup
	//	TblStmts map[string]*TableStatementGroup
}

func (dbi *DBInstance) Init(ip string, port int, user string, passwd string, dbname string, maxOpen int, maxIdle int) error {
	s := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, passwd, ip, port, dbname)
	db, err := sql.Open("mysql", s)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	if maxOpen > 0 {
		db.SetMaxOpenConns(maxOpen)
	}
	if maxIdle > 0 {
		db.SetMaxIdleConns(maxIdle)
	}
	dbi.MyDB = db
	dbi.closing = false
	//	dbi.requestQueue = make(chan *QueryRequest, 500)
	//	dbi.TblStmts = make(map[string]*TableStatementGroup)

	dbi.requestQueue = new(list.List)
	dbi.dealQueue = new(list.List)
	dbi.waitGroup.Add(1)
	go dbi.DealRequestInQueue()
	log.Infof("DBInstance Init %s,%d,%d", s, maxOpen, maxIdle)
	return nil
}

func (dbi *DBInstance) SetClose() {
	dbi.closing = true
	dbi.waitGroup.Wait()
}

func (dbi *DBInstance) Ping() {
	if dbi.MyDB != nil {
		dbi.QueryOneRow("select 1")
	}
}

func (dbi *DBInstance) StartTx() (*sql.Tx, error) {
	return dbi.MyDB.Begin()
}

func (dbi *DBInstance) Prepare(query string) (*sql.Stmt, error) {
	return dbi.MyDB.Prepare(query)
}

func (dbi *DBInstance) Exec(query string, args ...interface{}) (sql.Result, error) {
	//	log.Debugf("Exec: %v", query)
	result, err := dbi.MyDB.Exec(query, args...)
	if err != nil {
		log.Error(err)
	}
	return result, err
}
func (dbi *DBInstance) ExecNoWait(query string, args ...interface{}) {

	if dbi.closing {
		//服务器正在关闭时，同步执行完
		dbi.Exec(query, args...)
	} else {
		req := &QueryRequest{
			query: query,
			args:  args,
		}
		dbi.queueLock.Lock()
		dbi.requestQueue.PushBack(req)
		dbi.queueLock.Unlock()
	}
}

//func (dbm *DBMgr) QueryRow(query string, args ...interface{}) (*sql.Row, error) {
//	return dbm.QueryRow(query, args...)
//}

func (dbi *DBInstance) QueryOneRow(query string, args ...interface{}) (map[string]string, error) {
	rows, err := dbi.MyDB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	found := false
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
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

//查询多条数据
//limitRows 小于1时，返回全部数据
func (dbi *DBInstance) QueryRows(limitRows int, query string, args ...interface{}) ([]map[string]string, []string, error) {
	defer func() {
		if er := recover(); er != nil {
			fmt.Println("query %v err[%v]", query, er)
		}
	}()

	//	fmt.Println("DEBUG:query %v ", query)

	rows, err := dbi.MyDB.Query(query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	maxRows := limitRows
	if maxRows < 1 {
		maxRows = 100 //默认按100创建
	}

	retMap := make([]map[string]string, 0, maxRows)

	found := false
	rowCnt := 0
	// Fetch rows
	for rows.Next() {
		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, nil, err
		}
		found = true
		tmpMap := make(map[string]string)
		for i := 0; i < len(values); i++ {
			tmpMap[columns[i]] = string(values[i])
		}
		retMap = append(retMap, tmpMap)

		rowCnt++
		if limitRows > 0 && rowCnt >= limitRows {
			break
		}
	}
	if found {
		return retMap, columns, nil

	} else {
		return nil, columns, ErrorNoRow
	}
}

//func (dbi *DBInstance) RegisterStatement(tblName string, kind int, sqlstr string) (*sql.Stmt, error) {
//	stmt, err := dbi.Prepare(sqlstr)
//	if err != nil {
//		return nil, err
//	}
//	if group, ok := dbi.TblStmts[tblName]; ok {
//		group.Stmts[kind] = stmt
//	} else {
//		dbi.TblStmts[tblName] = &TableStatementGroup{
//			Stmts: make(map[int]*sql.Stmt),
//		}
//		dbi.TblStmts[tblName].Stmts[kind] = stmt
//	}
//	return stmt, nil
//}
//
//func (dbi *DBInstance) GetStatement(tblName string, kind int) *sql.Stmt {
//	if group, ok := dbi.TblStmts[tblName]; ok {
//		if st, yes := group.Stmts[kind]; yes {
//			return st
//		} else {
//			return nil
//		}
//	} else {
//		return nil
//	}
//}
//
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
				dbi.Exec(q.query, q.args...)
			}
			dbi.dealQueue.Init()
		} else {
			time.Sleep(time.Millisecond * 10)
		}

		if dbi.closing && dbi.requestQueue.Len() == 0 {
			return
		}
	}
}
