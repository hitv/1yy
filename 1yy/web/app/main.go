package main

import (
	"net/http"
	"runtime"

	"hi.tv/1yy/env"
	"hi.tv/1yy/models"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	conf := env.NewAppConfig("./app.conf")

	models.InitDB(conf.Dsn)
	env.InitDefaultApp(conf)

	router := InitRoute(conf.AssetPath)
	http.ListenAndServe(conf.Addr, router)
}
