package models

import (
	"bitbucket.org/chai2010/pwdgen/ini"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gocache "github.com/pmylund/go-cache"
	"sync"
	"time"
)

var (
	db        *sql.DB
	cache     *gocache.Cache
	posterFmt string
	staticFmt string
	pvChan    = make(chan int, 100000)
)

func initPVCounter() {
	lock := &sync.Mutex{}
	totalCounter := make(map[int]int)
	hourlyCounter := make(map[int]int)
	go func() {
		for vid := range pvChan {
			lock.Lock()
			totalCounter[vid]++
			hourlyCounter[vid]++
			lock.Unlock()
		}
	}()

	go func() {
		seq := 0
		totalStmt, err := db.Prepare("UPDATE video SET total_play_count=total_play_count+? WHERE id=?")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		hourlyStmt, err := db.Prepare("UPDATE video SET hourly_play_count=? WHERE id=?")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for {
			<-time.After(time.Minute * 5)
			lock.Lock()
			seq++
			for vid, count := range totalCounter {
				totalStmt.Exec(count, vid)
				totalCounter[vid] = 0
			}
			if seq == 12 {
				for vid, count := range hourlyCounter {
					hourlyStmt.Exec(count, vid)
					hourlyCounter[vid] = 0
				}
				seq = 0
			}
			lock.Unlock()
		}
	}()
}

//初始化
func Init(dict ini.Dict) error {
	cache = gocache.New(time.Minute*30, time.Minute*1)
	dsn, err := dict.GetString("common", "dsn")
	if err != nil {
		return err
	}

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	initPVCounter()
	return nil
}

func PVStatistics(vid int) {
	pvChan <- vid
}
func PosterUrl(poster string) string {
	config, err := GetConfig("PosterFormat")
	if err != nil {
		return ""
	}
	return fmt.Sprintf(config.Value, poster)
}
func StaticUrl(filePath string) string {
	config, err := GetConfig("StaticFormat")
	if err != nil {
		return ""
	}
	return fmt.Sprintf(config.Value, filePath)
}

func FlushModelCache() {
	cache.Flush()
}
