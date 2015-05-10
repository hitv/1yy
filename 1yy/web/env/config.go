package env

import (
	"log"
	"os"

	"github.com/robfig/config"

	"hi.tv/1yy/libs/render"
)

type Env string

const (
	EnvDev  Env = "development"
	EnvProd Env = "production"
)

func (e Env) IsDev() bool {
	return e == EnvDev
}

func (e Env) IsProd() bool {
	return e == EnvProd
}

type AppConfig struct {
	Environment         Env
	Addr                string
	Dsn                 string
	AssetPath           string
	AssetPrefix         string
	RedisHost           string
	RedisPassword       string
	RedisMaxIdle        int
	CacheDefaultExpires int
	SessionName         string
	SessionKeyPairs     []byte
	Logger              *log.Logger
	RenderOpt           render.Options
}

func NewAppConfig(confPath string) (appConf *AppConfig, err error) {
	conf, err := config.ReadDefault(confPath)
	if err != nil {
		return
	}

	environment, err := conf.String("", "environment")
	if err != nil {
		return
	}

	env := Env(environment)

	addr, err := conf.String("http", "addr")
	if err != nil {
		return
	}

	dsn, err := conf.String("mysql", "dsn")
	if err != nil {
		return
	}

	redisHost, err := conf.String("redis", "host")
	if err != nil {
		redisHost = ""
	}

	redisPassword, err := conf.String("redis", "password")
	if err != nil {
		redisPassword = ""
	}

	redisMaxIdle, err := conf.Int("redis", "max_idle")
	if err != nil {
		redisMaxIdle = 0
	}

	cacheExpires, err := conf.Int("cache", "default_expires")
	if err != nil {
		cacheExpires = 600
	}

	sessionKeyPairs, err := conf.String("session", "key_pairs")
	if err != nil {
		return
	}

	sessionName, err := conf.String("session", "name")
	if err != nil {
		return
	}

	assetPath, err := conf.String("path", "asset_path")
	if err != nil {
		assetPath = ""
	}

	assetPrefix, err := conf.String("path", "asset_prefix")
	if err != nil {
		assetPrefix = "/"
	}

	compressHTML := false
	if env.IsProd() {
		compressHTML = true
	}

	appConf = &AppConfig{
		Environment:         env,
		Addr:                addr,
		Dsn:                 dsn,
		AssetPath:           assetPath,
		AssetPrefix:         assetPrefix,
		RedisHost:           redisHost,
		RedisPassword:       redisPassword,
		RedisMaxIdle:        redisMaxIdle,
		CacheDefaultExpires: cacheExpires,
		SessionName:         sessionName,
		SessionKeyPairs:     []byte(sessionKeyPairs),
		Logger:              log.New(os.Stdout, "", log.LstdFlags),
		RenderOpt: render.Options{
			Layout:       "layout",
			CompressHTML: compressHTML,
		},
	}
	return
}
