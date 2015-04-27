package controllers

import "github.com/gocraft/web"

type Index struct {
	*Base
}

func (c *Index) Index(rw web.ResponseWriter, req *web.Request) {
	c.HTML(200, "index", nil)
}
