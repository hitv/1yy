package env

import (
	"html/template"
	"log"
	"path"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"hi.tv/1yy/libs/caches"
	"hi.tv/1yy/libs/render"
)

var DefaultApp *App

type App struct {
	Environment Env
	Logger      *log.Logger
	Config      *AppConfig
	Render      render.RenderHandler
	redisPool   *redis.Pool
	cache       caches.Cache
}

func templateFuncs(conf *AppConfig) []template.FuncMap {
	return []template.FuncMap{
		template.FuncMap{
			"asset": func(filePath string) string {
				if conf.Environment == EnvDev {
					filePath = strings.Replace(filePath, ".min.", ".src.", 1)
					//filePath += "?_t=" + strconv.FormatInt(time.Now().Unix(), 10)
				}
				return path.Clean(path.Join(conf.AssetPrefix, filePath))
			},
		},
	}
}

func prepareConfig(conf *AppConfig) {
	funcs := templateFuncs(conf)
	funcs = append(funcs, conf.RenderOpt.Funcs...)
	conf.RenderOpt.Funcs = funcs
}

func (a *App) RedisPool() *redis.Pool {
	if a.redisPool == nil {
		a.redisPool = caches.NewRedisPool(&caches.RedisConfig{
			Host:     a.Config.RedisHost,
			Password: a.Config.RedisPassword,
			MaxIdle:  a.Config.RedisMaxIdle,
		})
	}
	return a.redisPool
}

func (a *App) Cache() caches.Cache {
	if a.cache == nil {
		expires := time.Duration(a.Config.CacheDefaultExpires) * time.Second
		a.cache = caches.NewRedisCache(a.RedisPool(), expires)
	}
	return a.cache
}

func NewApp(conf *AppConfig) *App {
	prepareConfig(conf)
	return &App{
		Environment: conf.Environment,
		Logger:      conf.Logger,
		Config:      conf,
		Render:      render.NewRender(conf.RenderOpt),
	}
}

func InitDefaultApp(conf *AppConfig) {
	DefaultApp = NewApp(conf)
	return
}
