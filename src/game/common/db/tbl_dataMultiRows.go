package db

import (
	"errors"
	"fmt"
	util "game/common/utils"
	"strconv"
	"strings"
	"sync"

	log "github.com/cihub/seelog"
)

type SingleData struct {
	mu                        sync.RWMutex
	Data                      map[string]string
	DirtyProp                 map[string]string
	lastInsertIndex           int64
	setIndexBylastInsertIndex bool
}

func NewSingleData() *SingleData {
	return &SingleData{
		Data:                      make(map[string]string),
		DirtyProp:                 make(map[string]string),
		setIndexBylastInsertIndex: false,
	}
}

func (sd *SingleData) SetAutoIndex() {
	sd.setIndexBylastInsertIndex = true
}

func (sd *SingleData) GetInsertIndex() int64 {
	return sd.lastInsertIndex
}

func (sd *SingleData) SetPropValue(fieldName string, value string) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	if _, ok := sd.Data[fieldName]; ok {
		sd.Data[fieldName] = value
		sd.DirtyProp[fieldName] = value
	} else {
		log.Errorf("key not in map [%s]", fieldName)
	}
}

func (sd *SingleData) SetIntPropValue(fieldName string, value int32) {
	sd.SetPropValue(fieldName, strconv.FormatInt(int64(value), 10))
}

func (sd *SingleData) SetInt64PropValue(fieldName string, value int64) {
	sd.SetPropValue(fieldName, strconv.FormatInt(int64(value), 10))
}

func (sd *SingleData) GetValue(fieldKey string, defaultValue string) string {
	sd.mu.RLock()
	defer sd.mu.RUnlock()
	if s, ok := sd.Data[fieldKey]; ok {
		return s
	} else {
		return defaultValue
	}
}

func (sd *SingleData) GetPropValue(fieldKey string) string {
	return sd.GetValue(fieldKey, "")
}

func (sd *SingleData) GetIntPropValue(fieldKey string) int {
	value := sd.GetValue(fieldKey, "0")
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	} else {
		return v
	}
}

func (sd *SingleData) GetInt32PropValue(fieldKey string) int32 {
	value := sd.GetValue(fieldKey, "0")
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	} else {
		return int32(v)
	}
}

func (sd *SingleData) GetInt64PropValue(fieldKey string) int64 {
	value := sd.GetValue(fieldKey, "0")
	v, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	} else {
		return v
	}
}

func (sd *SingleData) InsertToDB(dbi *DBInstance, TblName string) error {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	var (
		keyStr   string = ""
		valueStr string = ""
	)
	for k, v := range sd.DirtyProp {
		keyStr = fmt.Sprintf("%s,`%s`", keyStr, k)
		valueStr = fmt.Sprintf("%s,\"%s\"", valueStr, util.EscapeSqlString(v))
	}
	sd.DirtyProp = make(map[string]string)

	//去掉开始的逗号
	keyStr = keyStr[1:]
	valueStr = valueStr[1:]

	q := fmt.Sprintf("insert into %s (%s) values (%s) ", TblName, keyStr, valueStr)
	log.Info(q)
	if !sd.setIndexBylastInsertIndex {
		dbi.ExecNoWait(q)
	} else {
		result, err := dbi.Exec(q)
		if err != nil {
			log.Errorf("sql[%v] err[%v]", q, err)
			return err
		}
		sd.lastInsertIndex, _ = result.LastInsertId()
	}

	return nil
}

func (sd *SingleData) DeleteFromDB(dbi *DBInstance, TblName string, keyNames []string, singleKey string) error {
	sd.mu.Lock()
	defer sd.mu.Unlock()

	keyStr := fmt.Sprintf(" %s=\"%s\"", singleKey, util.EscapeSqlString(sd.Data[singleKey]))
	for i := 0; i < len(keyNames); i++ {
		keyStr = fmt.Sprintf("%s and %s=\"%s\"", keyStr, keyNames[i], util.EscapeSqlString(sd.Data[keyNames[i]]))
	}

	if keyStr == "" {
		return fmt.Errorf("no data delete")
	}

	//	if strings.HasSuffix(keyStr, "and ") {
	//		fmt.Println(strings.TrimSuffix(keyStr, "and ")) //xiaowei
	//
	//	}
	keyStr = strings.TrimRight(keyStr, "and ")

	q := fmt.Sprintf("delete from  %s where %s ", TblName, keyStr)
	log.Info(q)
	dbi.ExecNoWait(q)

	return nil
}

func (sd *SingleData) SyncToDB(dbi *DBInstance, tblName string, keyNames []string, singleKey string) error {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	if len(sd.DirtyProp) > 0 {
		q := fmt.Sprintf("Update %s set ", tblName)
		for k, v := range sd.DirtyProp {
			q = fmt.Sprintf("%s %s=\"%s\",", q, k, util.EscapeSqlString(v))
		}
		q = q[:len(q)-1]
		sd.DirtyProp = make(map[string]string)

		where := fmt.Sprintf(" %s=\"%s\"", singleKey, util.EscapeSqlString(sd.Data[singleKey]))
		for i := 0; i < len(keyNames); i++ {
			where = fmt.Sprintf("%s and %s=\"%s\"", where, keyNames[i], util.EscapeSqlString(sd.Data[keyNames[i]]))
		}

		q = fmt.Sprintf("%s where %s", q, where)
		dbi.ExecNoWait(q)
	}
	return nil
}

type TblDataMulti struct {
	TblName          string
	KeyNames         []string
	Keys             []interface{}
	SingleKey        string
	Columns          []string
	propLock         sync.RWMutex
	PropsMap         map[string]*SingleData
	PropsArr         []*SingleData
	loaded           bool
	loading          bool
	StatementSqlLoad string
	specialCond      string
	limit            int
}

func CreateTblDataMulti(tblName string, keyNames []string, keys []interface{}, singleKey string, speciCon string, limit int) *TblDataMulti {
	if len(keyNames) != len(keys) {
		return nil
	}
	t := new(TblDataMulti)
	t.TblName = tblName
	t.KeyNames = keyNames
	t.Keys = keys
	t.SingleKey = singleKey
	t.PropsMap = make(map[string]*SingleData)
	if limit > 0 {
		t.PropsArr = make([]*SingleData, 0, limit*2)
	} else {
		t.PropsArr = make([]*SingleData, 0, 100)
	}
	t.loaded = false
	t.loading = false
	t.specialCond = speciCon
	t.limit = limit
	var s string = ""
	if len(keyNames) > 0 {
		s = " where "
	}
	for i := 0; i < len(keyNames); i++ {
		s = fmt.Sprintf("%s %s=? and ", s, t.KeyNames[i])
	}

	if len(keyNames) > 0 {
		s = strings.TrimRight(s, "and ")
	}

	s = fmt.Sprintf("%s %s", s, speciCon)

	if limit > 0 {
		s = fmt.Sprintf("%s limit %d", s, limit)
	}
	t.StatementSqlLoad = fmt.Sprintf("select * from %s %s", t.TblName, s)
	return t
}

func (tb *TblDataMulti) Load(dbi *DBInstance) (er error) {
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

	data, columns, err := dbi.QueryRows(tb.limit, tb.StatementSqlLoad, tb.Keys...)
	if err != nil && err != ErrorNoRow {
		return err
	}

	tb.Columns = columns

	if err != ErrorNoRow {
		for _, v := range data {
			sd := NewSingleData()
			sd.Data = v
			tb.PropsMap[sd.Data[tb.SingleKey]] = sd
			tb.PropsArr = append(tb.PropsArr, sd)
		}
	}

	tb.loaded = true
	return nil
}

func (tb *TblDataMulti) SetPropValue(key string, fieldName string, value string) {
	tb.propLock.RLock()
	defer tb.propLock.RUnlock()
	if v, ok := tb.PropsMap[key]; ok {
		v.SetPropValue(fieldName, value)
	} else {
		log.Infof("key not in map [%s]", key)
	}
}
func (tb *TblDataMulti) GetIntPropValue(mapKey string, fieldKey string) int {
	tb.propLock.RLock()
	defer tb.propLock.RUnlock()
	if v, ok := tb.PropsMap[mapKey]; ok {
		return v.GetIntPropValue(fieldKey)
	} else {
		return 0
	}
}
func (tb *TblDataMulti) GetInt64PropValue(mapKey string, fieldKey string) int64 {
	tb.propLock.RLock()
	defer tb.propLock.RUnlock()
	if v, ok := tb.PropsMap[mapKey]; ok {
		return v.GetInt64PropValue(fieldKey)
	} else {
		return 0
	}
}

func (tb *TblDataMulti) CreateNewRecord() *SingleData {
	sd := NewSingleData()
	for i := 0; i < len(tb.Columns); i++ {
		sd.Data[tb.Columns[i]] = ""
	}
	for i := 0; i < len(tb.KeyNames); i++ {
		sd.SetPropValue(tb.KeyNames[i], tb.Keys[i].(string))
	}
	return sd
}

func (tb *TblDataMulti) DeleteRecord(sd *SingleData, dbi *DBInstance) error {
	tb.propLock.Lock()
	defer tb.propLock.Unlock()
	err := sd.DeleteFromDB(dbi, tb.TblName, tb.KeyNames, tb.SingleKey)
	if err != nil {
		return err
	} else {
		delete(tb.PropsMap, sd.Data[tb.SingleKey])
		for k, v := range tb.PropsArr {
			if v == sd {
				tb.PropsArr = append(tb.PropsArr[:k], tb.PropsArr[k+1:]...)
				break
			}
		}
	}
	return nil
}

func (tb *TblDataMulti) UpdateRecord(sd *SingleData, dbi *DBInstance) error {
	tb.propLock.Lock()
	defer tb.propLock.Unlock()
	err := sd.SyncToDB(dbi, tb.TblName, tb.KeyNames, tb.SingleKey)
	if err != nil {
		return err
	} else {
		tb.PropsMap[sd.Data[tb.SingleKey]] = sd
		return nil
	}

}

func (tb *TblDataMulti) AddNewRecordAndDeleteOld(sd *SingleData, dbi *DBInstance, first bool, delFromDB bool) error {
	tb.propLock.Lock()
	defer tb.propLock.Unlock()
	if _, ok := tb.PropsMap[sd.Data[tb.SingleKey]]; ok {
		return errors.New("add duplicate record")
	}
	err := sd.InsertToDB(dbi, tb.TblName)
	if err != nil {
		log.Error(err)
		return err
	}

	if sd.setIndexBylastInsertIndex {
		sd.Data[tb.SingleKey] = strconv.FormatInt(sd.lastInsertIndex, 10)
		tb.PropsMap[sd.Data[tb.SingleKey]] = sd
	} else {
		tb.PropsMap[sd.Data[tb.SingleKey]] = sd
	}

	if first {
		newArr := []*SingleData{sd}
		tb.PropsArr = append(newArr, tb.PropsArr...)
	} else {
		tb.PropsArr = append(tb.PropsArr, sd)
	}

	if tb.limit > 0 && len(tb.PropsArr) > tb.limit {
		if first {
			de := tb.PropsArr[len(tb.PropsArr)-1]
			delete(tb.PropsMap, de.Data[tb.SingleKey])
			if delFromDB {
				err := de.DeleteFromDB(dbi, tb.TblName, tb.KeyNames, tb.SingleKey)
				if err != nil {
					return err
				}
			}
			tb.PropsArr = tb.PropsArr[:len(tb.PropsArr)-1]
		} else {
			de := tb.PropsArr[0]
			delete(tb.PropsMap, de.Data[tb.SingleKey])
			if delFromDB {
				err := de.DeleteFromDB(dbi, tb.TblName, tb.KeyNames, tb.SingleKey)
				if err != nil {
					return err
				}
			}
			tb.PropsArr = tb.PropsArr[1:]
		}
	}

	return nil
}

//添加新记录
//first : 新记录加在原记录前面还是后面
func (tb *TblDataMulti) AddNewRecord(sd *SingleData, dbi *DBInstance, first bool) error {
	return tb.AddNewRecordAndDeleteOld(sd, dbi, first, false)
}

func (tb *TblDataMulti) SyncToDB(dbi *DBInstance) error {
	tb.propLock.RLock()
	defer tb.propLock.RUnlock()
	for _, v := range tb.PropsMap {
		err := v.SyncToDB(dbi, tb.TblName, tb.KeyNames, tb.SingleKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tb *TblDataMulti) GetSingleData(key string) *SingleData {
	tb.propLock.RLock()
	defer tb.propLock.RUnlock()
	if v, ok := tb.PropsMap[key]; ok {
		return v
	} else {
		return nil
	}
}
