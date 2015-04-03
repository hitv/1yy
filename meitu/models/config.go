package models

import (
	"fmt"
	"time"
)

type MtConfig struct {
	Key    string //配置键
	Value  string //配置值
	Remark string //备注
}

func GetConfig(confKey string) (value string, err error) {
	key := fmt.Sprintf("Config/%s", confKey)
	data, exists := cache.Get(key)
	if exists {
		value = data.(string)
		return
	}
	row := db.QueryRow("SELECT `value` FROM `config` WHERE `key`=?", confKey)
	err = row.Scan(&value)
	if err != nil {
		return
	}
	cache.Set(key, value, time.Minute*5)
	return
}
