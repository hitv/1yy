package controllers

type Video struct {
	*Base
}

func (c *Video) List() {
	c.Render.HTML(200, "list", nil)
}

func (c *Video) Detail() {
	c.Render.HTML(200, "detail", map[string]interface{}{
		"params": c.Req.PathParams,
	})
}
