package main

import (
	"github.com/gocraft/web"
	"github.com/robfig/config"
	"hi.tv/web/app/controllers"
)

func InitRoute(conf *config.Config) *web.Router {
	root := web.New(controllers.Base{})
}
