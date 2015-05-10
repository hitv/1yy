package main

import (
	"log"
	"net/http"
	"runtime"

	"hi.tv/1yy/models"
	"hi.tv/1yy/web/env"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	conf, err := env.NewAppConfig("./app.conf")
	if err != nil {
		panic(err)
	}

	models.InitDB(conf.Dsn)
	env.InitDefaultApp(conf)

	router := InitRoute(conf.AssetPath)

	log.Printf("Start listen to %s\n", conf.Addr)
	err = http.ListenAndServe(conf.Addr, router)
	if err != nil {
		log.Printf("Start server error: %s\n", err)
	}
}
