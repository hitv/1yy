package types

import (
	"github.com/bitly/go-simplejson"
	"hi.tv/1yy/libs/mivideo/errors"
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
		err = errors.ErrFormatInvalid
		return
	}

	json = simplejson.NewFromInterface(channelBlocks[0])
	groupBlocks, err := json.Get("blocks").Array()
	if err != nil {
		err = errors.ErrFormatInvalid
		return
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
			continue
		}
		for _, item := range items {
			json := simplejson.NewFromInterface(item)
			groupItem, err := parseGroupItem(json)
			if err != nil || (groupItem.Entity != "svideo" && groupItem.Entity != "pvideo") {
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
		err = errors.ErrFormatInvalid
		return
	}

	json = simplejson.NewFromInterface(channelBlocks[1])
	title := json.Get("title").MustString()
	subTitle := json.Get("sub_titile").MustString()
	targetURL := json.GetPath("target", "url").MustString()

	channel = &Channel{
		Title:     title,
		SubTitle:  subTitle,
		TargetURL: targetURL,
		Groups:    groups,
	}

	return
}
