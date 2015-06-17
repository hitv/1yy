package controllers

import (
	"log"
	"time"

	"github.com/gofly/web"
	"hi.tv/1yy/libs/caches"
	"hi.tv/1yy/libs/render"
	"hi.tv/1yy/web/env"
)

var (
	ErrCacheMiss = caches.ErrCacheMiss
)

type Base struct {
	Rw     web.ResponseWriter `inject`
	Req    *web.Request       `inject`
	Render render.Render
	Cache  caches.Cache
	*env.App
}

func (c *Base) Init(next web.NextMiddlewareFunc) {
	c.App = env.DefaultApp
	c.Render = c.App.Render(c.Rw, c.Req.Request)
	c.Cache = c.App.Cache()
	next(c.Rw, c.Req)
}

func (c *Base) RequestLog(next web.NextMiddlewareFunc) {
	startTime := time.Now()
	next(c.Rw, c.Req)
	duration := time.Since(startTime)
	log.Printf("%s %s %d - %s\n", c.Req.Method, c.Req.URL.Path, c.Rw.StatusCode(), duration.String())
}
