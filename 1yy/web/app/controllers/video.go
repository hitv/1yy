package controllers

import "github.com/gocraft/web"

type Video struct {
	*Base
}

func (c *Video) List(rw web.ResponseWriter, req *web.Request) {
	c.HTML(200, "list", nil)
}
