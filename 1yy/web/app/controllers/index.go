package controllers

import (
	"time"

	"hi.tv/1yy/libs/mivideo"
)

type Index struct {
	*Base
}

func (c *Index) getHomeData() (homeData *mivideo.HomeData, err error) {
	homeData = &mivideo.HomeData{}
	err = c.Cache.Get("home-data", homeData)
	if err == nil {
		return
	}
	if err != ErrCacheMiss {
		return
	}
	miVideoService := mivideo.NewMiVideoService(
		c.App.Config.MiHost,
		c.App.Config.MiApi,
		c.App.Config.MiKey,
		c.App.Config.MiToken,
	)
	homeData, err = miVideoService.FetchHomeData()
	if err != nil {
		return
	}
	err = c.Cache.Set("home-data", homeData, time.Hour)
	return
}
func (c *Index) Index() {
	homeData, err := c.getHomeData()
	if err != nil {
		c.Logger.Printf("getHomeData error: %s", err)
		c.Render.Error(500)
		return
	}
	c.Render.HTML(200, "index", homeData.Hot)
}

func (c *Index) Best() {
	homeData, err := c.getHomeData()
	if err != nil {
		c.Logger.Printf("getHomeData error: %s", err)
		c.Render.Error(500)
		return
	}
	c.Render.JSON(200, homeData.Best)
}

func (c *Index) Rank() {
	homeData, err := c.getHomeData()
	if err != nil {
		c.Logger.Printf("getHomeData error: %s", err)
		c.Render.Error(500)
		return
	}
	c.Render.JSON(200, homeData.Rank)
}
