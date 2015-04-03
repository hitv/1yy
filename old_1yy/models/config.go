package models

import (
	"time"
)

var configCacheKey = "config"

type ConfigItem struct {
	Name   string
	Value  string
	Remark string
}

func GetConfig(key string) (*ConfigItem, error) {
	var configs map[string]*ConfigItem
	data, exists := cache.Get(configCacheKey)
	if exists {
		configs = data.(map[string]*ConfigItem)
	} else {
		configs = make(map[string]*ConfigItem)
		rows, err := db.Query("SELECT * FROM config")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			c := &ConfigItem{}
			err = rows.Scan(&c.Name, &c.Value, &c.Remark)
			if err != nil {
				return nil, err
			}
			configs[c.Name] = c
		}
		cache.Set(configCacheKey, configs, time.Minute*5)
	}
	return configs[key], nil
}

func FlushConfigCache() {
	cache.Delete(configCacheKey)
}
