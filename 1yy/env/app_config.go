package env

import (
	"log"

	"hi.tv/1yy/libs/render"
)

type Env string

const (
	EnvDev  Env = "development"
	EnvProd Env = "production"
)

type AppConfig struct {
	Environment         Env
	RedisHost           string
	RedisPassword       string
	RedisMaxIdle        int
	CacheDefaultExpires int
	SessionName         string
	SessionKeyPairs     []byte
	Logger              *log.Logger
	RenderOpt           render.Options
	AssetPrefix         string
}
