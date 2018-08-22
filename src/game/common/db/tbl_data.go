package db

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	util "game/common/utils"

	log "github.com/cihub/seelog"
)

type TblData struct {
	TblName          string
	KeyNames         []string
	Keys             []interface{}
	mu               sync.RWMutex
	props            map[string]string
	dirtyProp        map[string]string
	loaded           bool
	loading          bool
	StatementSqlLoad string
}

func CreateTblData(tblName string, keyNames []string, keys []interface{}) *TblData {
	if len(keyNames) != len(keys) {
		return nil
	}
	t := new(TblData)
	t.TblName = tblName
	t.KeyNames = keyNames
	t.Keys = keys
	t.props = make(map[string]string)
	t.dirtyProp = make(map[string]string)
	t.loaded = false
	t.loading = false
	var s string = ""
	for i := 0; i < len(keyNames); i++ {
		s = fmt.Sprintf("%s %s=? and", s, t.KeyNames[i])
	}
	s = s[:len(s)-4]
	t.StatementSqlLoad = fmt.Sprintf("select * from %s where %s", t.TblName, s)

	return t
}

func (tb *TblData) Load(dbi *DBInstance) (er error) {
	if tb.loaded || tb.loading {
		return errors.New("Tbl Loaded or Loading")
	}
	if dbi == nil {
		return errors.New("DBInstance nil")
	}
	tb.loading = true
	defer func() {
		if err := recover(); err != nil {
			er = errors.New(fmt.Sprintf("load %v err[%v]", tb.TblName, err))
		}
		tb.loading = false
	}()

	data, err := dbi.QueryOneRow(tb.StatementSqlLoad, tb.Keys...)
	if err != nil && err != ErrorNoRow {
		return err
	}
	if err == ErrorNoRow {
		tb.props = data
		for i := 0; i < len(tb.KeyNames); i++ {
			tb.props[tb.KeyNames[i]] = tb.Keys[i].(string)
		}
		return err
	}

	tb.props = data
	tb.loaded = true
	return nil
}

func (tb *TblData) GetFields() []string {
	tb.mu.RLock()
	defer tb.mu.RUnlock()
	ret := make([]string, 0, len(tb.props))
	for k, _ := range tb.props {
		ret = append(ret, k)
	}
	return ret
}

func (tb *TblData) SetPropValue(fieldName string, value string) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if _, ok := tb.props[fieldName]; ok {
		tb.props[fieldName] = value
		tb.dirtyProp[fieldName] = value
	} else {
		log.Errorf("key not in map [%s]", fieldName)
	}
}

func (tb *TblData) GetValue(fieldName string, defaultValue string) string {
	tb.mu.RLock()
	defer tb.mu.RUnlock()
	if v, ok := tb.props[fieldName]; ok {
		return v
	} else {
		return defaultValue
	}
}
func (tb *TblData) GetPropValue(fieldName string) string {
	return tb.GetValue(fieldName, "")
}
func (tb *TblData) GetIntPropValue(key string) int {
	value := tb.GetValue(key, "0")
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	} else {
		return v
	}
}
func (tb *TblData) GetInt64PropValue(key string) int64 {
	value := tb.GetValue(key, "0")
	v, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	} else {
		return v
	}
}

func (tb *TblData) InsertToDB(dbi *DBInstance, nowait bool) error {
	var (
		keyStr   string = ""
		valueStr string = ""
	)
	tb.mu.RLock()
	defer tb.mu.RUnlock()
	for k, v := range tb.dirtyProp {
		keyStr = fmt.Sprintf("%s,`%s`", keyStr, k)
		valueStr = fmt.Sprintf("%s,\"%s\"", valueStr, util.EscapeSqlString(v))
	}
	for i := 0; i < len(tb.KeyNames); i++ {
		if _, ok := tb.dirtyProp[tb.KeyNames[i]]; !ok {

			keyStr = fmt.Sprintf("%s,%s", keyStr, tb.KeyNames[i])
			valueStr = fmt.Sprintf("%s,\"%s\"", valueStr, util.EscapeSqlString(tb.Keys[i].(string)))
		}
	}
	keyStr = keyStr[1:] //去掉逗号
	valueStr = valueStr[1:]

	q := fmt.Sprintf("insert into %s (%s) values (%s) ", tb.TblName, keyStr, valueStr)
	log.Info(q)
	if nowait {
		dbi.ExecNoWait(q)
	} else {
		_, err := dbi.Exec(q)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (tb *TblData) SyncToDB(dbi *DBInstance) error {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if len(tb.dirtyProp) > 0 {
		q := fmt.Sprintf("Update %s set ", tb.TblName)
		for k, v := range tb.dirtyProp {
			q = fmt.Sprintf("%s %s=\"%s\",", q, k, util.EscapeSqlString(v))
		}
		tb.dirtyProp = make(map[string]string)

		q = q[:len(q)-1]

		where := ""
		for i := 0; i < len(tb.KeyNames); i++ {
			where = fmt.Sprintf("%s %s=\"%s\" and ", where, tb.KeyNames[i], util.EscapeSqlString(tb.Keys[i].(string)))
		}
		where = where[:len(where)-4]

		//		log.Info(q)
		q = fmt.Sprintf("%s where %s", q, where)
		dbi.ExecNoWait(q)
	}
	return nil
}
