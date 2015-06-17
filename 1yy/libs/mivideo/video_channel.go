package mivideo

import "github.com/bitly/go-simplejson"

type VideoChannel struct {
	Title    string
	Slides   []*GroupItem
	Channels []*Channel
	*Time
}

func parseSlides(json *simplejson.Json) ([]*GroupItem, error) {
	items, err := json.Get("items").Array()
	if err == nil {
		slides := make([]*GroupItem, 0, len(items))
		for _, item := range items {
			groupItem, err := parseGroupItem(simplejson.NewFromInterface(item))
			if err != nil {
				return nil, err
			}
			if groupItem.Entity != "pvideo" {
				continue
			}
			slides = append(slides, groupItem)
		}
	} else {
		err = ErrFormatInvalid
	}
	return nil, err
}

func parseVideoChannel(json *simplejson.Json) (*VideoChannel, error) {
	title := json.Get("title").MustString()
	blocks, err := json.Get("blocks").Array()
	if err != nil {
		return nil, ErrFormatInvalid
	}

	times := parseTime(json.Get("times"))

	// slides
	var slides []*GroupItem
	if len(blocks) > 0 {
		json := simplejson.NewFromInterface(blocks[0])
		slides, err = parseSlides(json)
		if err != nil {
			//return nil, err
		}
	}

	channels := make([]*Channel, 0, len(blocks))
	for _, block := range blocks {
		json := simplejson.NewFromInterface(block)
		channel, err := parseChannel(json)
		if err != nil {
			//return nil, err
			continue
		}
		channels = append(channels, channel)
	}

	return &VideoChannel{
		Title:    title,
		Slides:   slides,
		Channels: channels,
		Time:     times,
	}, nil
}
