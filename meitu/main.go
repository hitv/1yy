package main

import (
	"bitbucket.org/chai2010/pwdgen/ini"
	"flag"
	"github.com/hoisie/web"
	"meitu/controllers"
	"meitu/models"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	confFile := flag.String("conf", "./config.ini", "配置文件路径")
	flag.Parse()

	dict, err := ini.LoadFile(*confFile)
	if err != nil {
		panic(err)
	}

	port, err := dict.GetString("common", "listen")
	if err != nil {
		panic(err)
	}

	dsn, err := dict.GetString("common", "dsn")
	if err != nil {
		panic(err)
	}

	err = models.InitModel(dsn)
	if err != nil {
		panic(err)
	}

	web.Get(`^/c/(\d+)\.html`, controllers.CategoryPage)
	web.Get(`^/albums/(\d+)/(\d+)\.json$`, controllers.CategoryAlbumsApi)
	web.Get(`^/a/(\d+)\.html$`, controllers.AlbumPage)
	web.Get(`^/images/(\d+)\.json$`, controllers.AlbumImagesApi)
	web.Run(port)
}
