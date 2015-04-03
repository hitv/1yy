package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DefaultDB   *sql.DB
	ErrNotFound = sql.ErrNoRows
)

func InitDB(dsn string) (err error) {
	DefaultDB, err = sql.Open("mysql", dsn)
	gorm.
	return
}
