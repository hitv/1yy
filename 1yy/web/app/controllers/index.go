package controllers

import (
	"fmt"

	"github.com/gocraft/web"
)

type Index struct {
	*Base
}

func (c *Index) Index(rw web.ResponseWriter, req *web.Request) {
	fmt.Printf("%v\n", &(c.Base))
	c.HTML(200, "index", nil)
}

func (c *Index) Test(rw web.ResponseWriter, req *web.Request) {
	c.HTML(200, "test", nil)
}
