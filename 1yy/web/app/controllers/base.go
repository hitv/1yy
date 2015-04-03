package controllers

import (
	"github.com/gocraft/web"
	"hi.tv/1yy/web/app/env"
)

type Base struct {
	App *env.App
}

func (c *Base) Init(rw *web.ResponseWriter, req *web.Request) {
	c.App = env.DefaultApp
}
