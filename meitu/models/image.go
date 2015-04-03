package models

import (
	"fmt"
)

type MtImage struct {
	Aid    int64  //所属图册id
	Sn     int64  //所在图册中的序号
	Url    string //图片URL
	Width  int64  //图片宽度
	Height int64  //图片高度
}

func GetAlbumImages(albumId int) (images []*MtImage, err error) {
	key := fmt.Sprintf("AlbumImages/%d", albumId)
	data, exists := cache.Get(key)
	if exists {
		images = data.([]*MtImage)
		return
	}
	rows, err := db.Query("SELECT aid,sn,url,width,height FROM image WHERE aid=? AND status='NORMAL' ORDER BY sn ASC", albumId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		image := &MtImage{}
		err = rows.Scan(&image.Aid, &image.Sn, &image.Url, &image.Width, &image.Height)
		if err != nil {
			return
		}
		images = append(images, image)
	}
	cache.Set(key, images, 0)
	return
}
