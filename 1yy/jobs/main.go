package main

import (
	"fmt"

	"github.com/hitv/1yy/models"
)

func main() {
	err := models.InitDB("root:root@tcp/1yingyuan?charset=utf8")
	if err != nil {
		fmt.Printf("models.InitDB error: %s", err)
		return
	}
	/*
		info, err := fetch.FetchChannelInfo(-1)
		if err != nil {
			fmt.Printf("FetchChannelInfo(-1) error: %s\n", err)
			return
		}
		err = info.Save()
		if err != nil {
			fmt.Printf("info.Save() error: %s\n", err)
			return
		}
	*/
	/*data, err := fetch.FetchMediaDetail(1075038)
	if err != nil {
		fmt.Printf("fetch.FetchMediaDetail(%d) error: %s", 1075038, err)
		return
	}
	fmt.Printf("data: %#v\n\n", data)
	*/
	channels, err := models.ChannelDao.GetAllFilter()
	if err != nil {
		fmt.Printf("ChannelDao.GetAllFilter() error: %s", err)
		return
	}
	for _, channel := range channels {
		var (
			page     = 1
			pageSize = 100
		)
		for {
			fmt.Printf("Start fetch channelId: %d, page: %d\n", channel.MiId, page)
			data, err := mivideo.FetchMediaFilter(channel.MiId, pageSize, page)
			if err != nil {
				fmt.Printf("mivideo.FetchMediaFilter(%d, %d, %d) error: %s", channel.MiId, pageSize, page, err)
				return
			}
			for _, video := range data.Data {
				id, err := video.Save()
				if err != nil {
					fmt.Printf("video.Save() error: %s\n", err)
					continue
				}
				fmt.Printf("video.Save() vid: %d\n", id)
			}
			if data.Count <= page*100 {
				break
			}
			page++
		}
	}
}
