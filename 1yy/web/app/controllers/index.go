package controllers

import (
	"github.com/gocraft/web"
	"hi.tv/1yy/libs/iqiyi"
)

type Index struct {
	*Base
}

func (c *Index) Index(rw web.ResponseWriter, req *web.Request) {
	str, err := iqiyi.GetM3u8("http://www.iqiyi.com/v_19rro1vd70.html")
	if err != nil {
		c.JSON(500, "GetM3u8 error:"+err.Error())
		return
	}

	c.JSON(200, str)
}
