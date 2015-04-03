package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gocache "github.com/pmylund/go-cache"
	"time"
)

var (
	db    *sql.DB
	cache *gocache.Cache
)

func InitModel(dsn string) (err error) {
	cache = gocache.New(time.Minute*30, time.Minute*5)
	db, err = sql.Open("mysql", dsn)
	return
}
