package server

import (
	log "github.com/cihub/seelog"
)

type Stroke struct {
	UID       int64
	Send      string
	Tel       string
	Content   string
	Timestamp int64
}

func CreateStroke(s Stroke) error {
	sql := "insert into stroke (Send,Tel,Content,Timestamp) value (?,?,?,?)"
	dbi := DBMgr.GetInstance()
	tx, _ := dbi.StartTx()
	_, err := tx.Exec("SET NAMES utf8mb4")
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(sql, s.Send, s.Tel, s.Content, s.Timestamp)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
