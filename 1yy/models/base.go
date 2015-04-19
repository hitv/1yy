package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	db          *xorm.Engine
	ErrNotFound = sql.ErrNoRows
)

func InitDB(dsn string) (err error) {
	db, err = xorm.NewEngine("mysql", dsn)
	return
}
