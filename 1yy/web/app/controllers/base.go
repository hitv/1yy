package controllers

import (
	"github.com/gocraft/web"
	"hi.tv/1yy/env"
	"hi.tv/1yy/libs/caches"
	"hi.tv/1yy/libs/render"
)

type Base struct {
	render.Render
	*env.App
	Cache caches.Cache
}

func (c *Base) Init(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.App = env.DefaultApp
	c.Cache = c.App.Cache()
	c.Render = c.App.Render(rw, req.Request)
	next(rw, req)
}
