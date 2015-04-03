package models

import (
	"fmt"
)

var (
	typeNameMap = map[string]string{
		"bdhd": "百度影音",
		"qvod": "快播",
	}
)

type TypeUrl struct {
	Name  string
	Code  string
	Slice []*Url
	Map   map[int64]*Url
}

type TypeUrls struct {
	Slice []*TypeUrl
	Map   map[string]*TypeUrl
}

type Url struct {
	Vid          int64
	Sn           int64
	Title        string
	Type         string
	Url          string
	ModifiedTime string
	CreateTime   string
}

func (url *Url) IsCurrent(urlType string, sn int64) bool {
	return url.Type == urlType && url.Sn == sn
}
func GetVideoTypeUrlMap(vid int) (map[string]*TypeUrl, error) {
	key := fmt.Sprintf("VideoTypeUrlMap/%d", vid)
	data, exists := cache.Get(key)
	var typeUrlMap map[string]*TypeUrl
	if exists {
		typeUrlMap = data.(map[string]*TypeUrl)
	} else {
		typeUrlMap = make(map[string]*TypeUrl)
		rows, err := db.Query("SELECT vid,sn,title,type,url FROM `url` WHERE `vid`=? AND `status`='NORMAL' ORDER BY sn DESC", vid)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			url := &Url{}
			err = rows.Scan(&url.Vid, &url.Sn, &url.Title, &url.Type, &url.Url)
			if err != nil {
				return nil, err
			}
			typeUrl := typeUrlMap[url.Type]
			if typeUrl == nil {
				typeUrl = &TypeUrl{
					Name:  typeNameMap[url.Type],
					Code:  url.Type,
					Slice: make([]*Url, 0),
					Map:   make(map[int64]*Url),
				}
				typeUrlMap[url.Type] = typeUrl
			}
			typeUrl.Slice = append(typeUrl.Slice, url)
			typeUrl.Map[url.Sn] = url
		}

		cache.Set(key, typeUrlMap, 0)
	}
	return typeUrlMap, nil
}

func GetSortedVideoTypeUrls(vid int, topType string) (*TypeUrls, error) {
	key := fmt.Sprintf("SortedVideoTypeUrls/%d/%s", vid, topType)
	var urls *TypeUrls
	data, exists := cache.Get(key)
	if exists {
		urls = data.(*TypeUrls)
	} else {
		typeUrlMap, err := GetVideoTypeUrlMap(vid)
		if err != nil {
			return nil, err
		}
		urls = &TypeUrls{
			Slice: make([]*TypeUrl, 0, len(typeNameMap)),
			Map:   make(map[string]*TypeUrl),
		}

		topTypeUrl := typeUrlMap[topType]
		if topTypeUrl != nil {
			urls.Slice = append(urls.Slice, topTypeUrl)
		}
		for urlType, typeUrl := range typeUrlMap {
			if urlType != topType {
				urls.Slice = append(urls.Slice, typeUrl)
			}
			urls.Map[urlType] = typeUrl
		}
		cache.Set(key, urls, 0)
	}
	return urls, nil
}

func GetVideoTypeUrls(vid uint64, urlType string) ([]*Url, error) {
	urls := make([]*Url, 0, 2000)
	rows, err := db.Query("SELECT vid,type,sn,title,url,create_time,modified_time FROM url WHERE vid=? AND type=?", vid, urlType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		u := &Url{}
		err = rows.Scan(&u.Vid, &u.Type, &u.Sn, &u.Title, &u.Url, &u.CreateTime, &u.ModifiedTime)
		urls = append(urls, u)
	}
	return urls, nil
}
