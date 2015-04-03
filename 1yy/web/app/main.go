package main

import (
	"time"

	"github.com/robfig/config"
	"hi.tv/1yy/models"
	"hi.tv/1yy/services/cache"
)

func parseConfig() {

}

func initDB(conf *config.Config) {
	dsn, err := conf.String("mysql", "dsn")
	if err != nil {
		panic("Can't find mysql->dsn in config file")
	}
	err = models.InitDB(dsn)
	if err != nil {
		panic(err)
	}
}

func RedisCacheService(conf *config.Config) cache.Cache {
	redisHost, err := conf.String("redis", "host")
	if err != nil {
		panic("Can't find redis->host in config file")
	}
	redisPool := cache.NewRedisPool(&cache.RedisConfig{
		Host: redisHost,
	})
	return cache.NewRedisCache(redisPool, time.Hour*24*7)
}

func main() {
	conf, err := config.ReadDefault("app.ini")
	if err != nil {
		panic(err)
	}
	initDB(conf)
	redisCache := RedisCacheService(conf)
}
