package mivideo

import (
	"github.com/bitly/go-simplejson"
)

type Channel struct {
	Title     string
	SubTitle  string
	TargetURL string
	Groups    []*Group
}

func parseChannel(json *simplejson.Json) (channel *Channel, err error) {
	channelBlocks, err := json.Get("blocks").Array()
	if err != nil || len(channelBlocks) < 2 {
		return nil, ErrFormatInvalid
	}

	json = simplejson.NewFromInterface(channelBlocks[0])
	groupBlocks, err := json.Get("blocks").Array()
	if err != nil {
		return nil, ErrFormatInvalid
	}

	var groups []*Group
	for _, groupBlock := range groupBlocks {
		var (
			groupItems []*GroupItem
			json       = simplejson.NewFromInterface(groupBlock)
		)
		title := json.Get("title").MustString()
		subTitle := json.Get("sub_title").MustString()
		targetURL := json.GetPath("target", "url").MustString()
		items, err := json.Get("items").Array()
		if err != nil {
			return nil, ErrFormatInvalid
		}
		for _, item := range items {
			json := simplejson.NewFromInterface(item)
			groupItem, err := parseGroupItem(json)
			if err != nil || groupItem.Entity != "svideo" && groupItem.Entity != "pvideo" {
				continue
			}
			groupItems = append(groupItems, groupItem)
		}
		if len(groupItems) == 0 {
			continue
		}

		groups = append(groups, &Group{
			Title:     title,
			SubTitle:  subTitle,
			TargetURL: targetURL,
			Items:     groupItems,
		})
	}

	if len(groups) == 0 {
		return nil, ErrFormatInvalid
	}

	json = simplejson.NewFromInterface(channelBlocks[1])
	title := json.Get("title").MustString()
	subTitle := json.Get("sub_titile").MustString()
	targetURL := json.GetPath("target", "url").MustString()

	return &Channel{
		Title:     title,
		SubTitle:  subTitle,
		TargetURL: targetURL,
		Groups:    groups,
	}, nil
}
