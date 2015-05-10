package controllers

type Channel struct {
	*Base
}

func (c *Base) All() {
	c.Render.JSON(200, c.Req.PathParams)
}

func (c *Base) List() {
	c.Render.JSON(200, c.Req.PathParams)
}

func (c *Base) Best() {
	c.Render.JSON(200, c.Req.PathParams)
}

func (c *Base) Category() {
	c.Render.JSON(200, c.Req.PathParams)
}
