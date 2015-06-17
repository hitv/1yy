package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	db          *xorm.Engine
	ErrNotExist = xorm.ErrNotExist
)

func InitDB(dsn string) (err error) {
	db, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Sync2(&RecGroupModel{}, &RecGroupItemModel{})
	return
}
